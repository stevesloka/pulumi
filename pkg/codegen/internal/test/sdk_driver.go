package test

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
)

// Defines an extra check logic that accepts the directory with the
// generated code, typically `$TestDir/$test.Directory/$language`.
type CodegenCheck func(t *testing.T, codedir string)

type sdkTest struct {
	Directory   string
	Description string

	// Extra checks for this test. They keys of this map
	// are of the form "$language/$check" such as "go/compile".
	Checks map[string]CodegenCheck

	// Skip checks, identified by "$language/$check".
	Skip codegen.StringSet

	// Do not compile the generated code for the languages in this set.
	// This is a helper form of `Skip`.
	SkipCompileCheck codegen.StringSet
}

const (
	// python = "python"
	nodejs = "nodejs"
	dotnet = "dotnet"
	golang = "go"
)

var sdkTests = []sdkTest{
	{
		Directory:   "input-collision",
		Description: "Schema with types that could potentially produce collisions (go).",
	},
	{
		Directory:   "dash-named-schema",
		Description: "Simple schema with a two part name (foo-bar)",
	},
	{
		Directory:        "external-resource-schema",
		Description:      "External resource schema",
		SkipCompileCheck: codegen.NewStringSet(nodejs, golang),
	},
	{
		Directory:        "nested-module",
		Description:      "Nested module",
		SkipCompileCheck: codegen.NewStringSet(dotnet, nodejs),
	},
	{
		Directory:        "nested-module-thirdparty",
		Description:      "Third-party nested module",
		SkipCompileCheck: codegen.NewStringSet(dotnet, nodejs),
	},
	{
		Directory:   "plain-schema-gh6957",
		Description: "Repro for #6957",
	},
	{
		Directory:   "resource-args-python",
		Description: "Resource args with same named resource and type",
	},
	{
		Directory:   "simple-enum-schema",
		Description: "Simple schema with enum types",
	},
	{
		Directory:   "simple-plain-schema",
		Description: "Simple schema with plain properties",
	},
	{
		Directory:   "simple-plain-schema-with-root-package",
		Description: "Simple schema with root package set",
	},
	{
		Directory:   "simple-resource-schema",
		Description: "Simple schema with local resource properties",
	},
	{
		Directory:   "simple-resource-schema-custom-pypackage-name",
		Description: "Simple schema with local resource properties and custom Python package name",
	},
	{
		Directory:        "simple-methods-schema",
		Description:      "Simple schema with methods",
		SkipCompileCheck: codegen.NewStringSet(nodejs, dotnet, golang),
	},
	{
		Directory:   "simple-yaml-schema",
		Description: "Simple schema encoded using YAML",
	},
	{
		Directory:        "provider-config-schema",
		Description:      "Simple provider config schema",
		SkipCompileCheck: codegen.NewStringSet(dotnet),
	},
	{
		Directory:        "replace-on-change",
		Description:      "Simple use of replaceOnChange in schema",
		SkipCompileCheck: codegen.NewStringSet(golang),
		Skip:             codegen.NewStringSet("nodejs/test"),
	},
	{
		Directory:        "resource-property-overlap",
		Description:      "A resource with the same name as it's property",
		SkipCompileCheck: codegen.NewStringSet(dotnet, nodejs),
	},
}

type SDKCodegenOptions struct {
	// Name of the programming language.
	Language string

	// Language-aware code generator; such as `GeneratePackage`.
	// from `codgen/dotnet`.
	GenPackage GenPkgSignature

	// Extra checks for all the tests. They keys of this map are
	// of the form "$language/$check" such as "go/compile".
	Checks map[string]CodegenCheck
}

// TestSDKCodegen runs the complete set of SDK code generation tests against a particular language's code
// generator. It also verifies that the generated code is structurally sound.
//
// An SDK code generation test consists of a schema and a set of expected outputs for each language. Each test is
// structured as a directory that contains that information:
//
//     test-directory/
//         schema.(json|yaml)
//         language-0
//         ...
//         language-n
//
// The schema is the only piece that must be manually authored. Once the schema has been written, the expected outputs
// can be generated by running `PULUMI_ACCEPT=true go test ./..." from the `pkg/codegen` directory.
func TestSDKCodegen(t *testing.T, opts *SDKCodegenOptions) {
	testDir := filepath.Join("..", "internal", "test", "testdata")

	for _, tt := range sdkTests {
		t.Run(tt.Description, func(t *testing.T) {
			dirPath := filepath.Join(testDir, filepath.FromSlash(tt.Directory))

			schemaPath := filepath.Join(dirPath, "schema.json")
			if _, err := os.Stat(schemaPath); err != nil && os.IsNotExist(err) {
				schemaPath = filepath.Join(dirPath, "schema.yaml")
			}

			files, err := GeneratePackageFilesFromSchema(schemaPath, opts.GenPackage)
			require.NoError(t, err)

			if !RewriteFilesWhenPulumiAccept(t, dirPath, opts.Language, files) {
				expectedFiles, err := LoadBaseline(dirPath, opts.Language)
				require.NoError(t, err)

				if !ValidateFileEquality(t, files, expectedFiles) {
					t.Fail()
				}
			}

			CopyExtraFiles(t, dirPath, opts.Language)

			// Merge language-specific global and
			// test-specific checks, with test-specific
			// having precedence.
			allChecks := make(map[string]CodegenCheck)
			for k, v := range opts.Checks {
				allChecks[k] = v
			}
			for k, v := range tt.Checks {
				allChecks[k] = v
			}

			// Define check filter.
			shouldSkipCheck := func(check string) bool {
				return false
			}

			codeDir := filepath.Join(testDir, options.Language)

			if hooks, ok := tt.HooksByLanguage[options.Language]; ok {
				runOne := func(hook SdkTestHook) {
					t.Run(hook.Name, func(t *testing.T) {
						hook.RunHook(&SdkTestEnv{
							T:       t,
							CodeDir: codeDir,
						})
					})
				}

				for _, hook := range hooks {
					runOne(hook)
				}
			}

			// Define check filter.
			shouldSkipCheck := func(check string) bool {

				// Only language-specific checks.
				if !strings.HasPrefix(check, opts.Language+"/") {
					return true
				}

				// Obey SkipCompileCheck to skip compile and test targets.
				if tt.SkipCompileCheck != nil &&
					tt.SkipCompileCheck.Has(opts.Language) &&
					(check == fmt.Sprintf("%s/compile", opts.Language) ||
						check == fmt.Sprintf("%s/test", opts.Language)) {
					return true
				}

				// Obey Skip.
				if tt.Skip != nil && tt.Skip.Has(check) {
					return true
				}

				return false
			}

			// Sort the checks in alphabetical order.
			var checkOrder []string
			for check := range allChecks {
				checkOrder = append(checkOrder, check)
			}
			sort.Strings(checkOrder)

			codeDir := filepath.Join(dirPath, opts.Language)

			// Perform the checks.
			for _, checkVar := range checkOrder {
				check := checkVar
				t.Run(check, func(t *testing.T) {
					if shouldSkipCheck(check) {
						t.Skip()
					}
					checkFun := allChecks[check]
					checkFun(t, codeDir)
				})
			}
		})
	}
}
