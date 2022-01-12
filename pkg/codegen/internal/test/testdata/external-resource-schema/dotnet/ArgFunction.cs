// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example
{
    public static class ArgFunction
    {
        public static Task<ArgFunctionResult> InvokeAsync(ArgFunctionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ArgFunctionResult>("example::argFunction", args ?? new ArgFunctionArgs(), options.WithVersion());

        public static Output<ArgFunctionResult> Invoke(ArgFunctionInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<ArgFunctionResult>("example::argFunction", args ?? new ArgFunctionInvokeArgs(), options.WithVersion());
    }


    public sealed class ArgFunctionArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public Pulumi.Random.RandomPet? Name { get; set; }

        public ArgFunctionArgs()
        {
        }
    }

    public sealed class ArgFunctionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<Pulumi.Random.RandomPet>? Name { get; set; }

        public ArgFunctionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class ArgFunctionResult
    {
        public readonly int? Age;

        [OutputConstructor]
        private ArgFunctionResult(int? age)
        {
            Age = age;
        }
    }
}
