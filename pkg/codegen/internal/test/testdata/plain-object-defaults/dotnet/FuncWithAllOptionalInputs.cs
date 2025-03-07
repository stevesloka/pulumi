// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi.Utilities;

namespace Pulumi.Example
{
    public static class FuncWithAllOptionalInputs
    {
        /// <summary>
        /// Check codegen of functions with all optional inputs.
        /// </summary>
        public static Task<FuncWithAllOptionalInputsResult> InvokeAsync(FuncWithAllOptionalInputsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<FuncWithAllOptionalInputsResult>("example::funcWithAllOptionalInputs", args ?? new FuncWithAllOptionalInputsArgs(), options.WithDefaults());

        /// <summary>
        /// Check codegen of functions with all optional inputs.
        /// </summary>
        public static Output<FuncWithAllOptionalInputsResult> Invoke(FuncWithAllOptionalInputsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<FuncWithAllOptionalInputsResult>("example::funcWithAllOptionalInputs", args ?? new FuncWithAllOptionalInputsInvokeArgs(), options.WithDefaults());
    }


    public sealed class FuncWithAllOptionalInputsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Property A
        /// </summary>
        [Input("a")]
        public Inputs.HelmReleaseSettings? A { get; set; }

        /// <summary>
        /// Property B
        /// </summary>
        [Input("b")]
        public string? B { get; set; }

        public FuncWithAllOptionalInputsArgs()
        {
            B = "defValue";
        }
    }

    public sealed class FuncWithAllOptionalInputsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Property A
        /// </summary>
        [Input("a")]
        public Input<Inputs.HelmReleaseSettingsArgs>? A { get; set; }

        /// <summary>
        /// Property B
        /// </summary>
        [Input("b")]
        public Input<string>? B { get; set; }

        public FuncWithAllOptionalInputsInvokeArgs()
        {
            B = "defValue";
        }
    }


    [OutputType]
    public sealed class FuncWithAllOptionalInputsResult
    {
        public readonly string R;

        [OutputConstructor]
        private FuncWithAllOptionalInputsResult(string r)
        {
            R = r;
        }
    }
}
