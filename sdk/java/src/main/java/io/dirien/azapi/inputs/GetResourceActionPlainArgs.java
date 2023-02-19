// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.azapi.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetResourceActionPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetResourceActionPlainArgs Empty = new GetResourceActionPlainArgs();

    /**
     * The name of the resource action. It&#39;s also possible to make Http requests towards the resource ID if leave this field empty.
     * 
     */
    @Import(name="action")
    private @Nullable String action;

    /**
     * @return The name of the resource action. It&#39;s also possible to make Http requests towards the resource ID if leave this field empty.
     * 
     */
    public Optional<String> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * A JSON object that contains the request body.
     * 
     */
    @Import(name="body")
    private @Nullable String body;

    /**
     * @return A JSON object that contains the request body.
     * 
     */
    public Optional<String> body() {
        return Optional.ofNullable(this.body);
    }

    /**
     * Specifies the Http method of the azure resource action. Allowed values are `POST` and `GET`. Defaults to `POST`.
     * 
     */
    @Import(name="method")
    private @Nullable String method;

    /**
     * @return Specifies the Http method of the azure resource action. Allowed values are `POST` and `GET`. Defaults to `POST`.
     * 
     */
    public Optional<String> method() {
        return Optional.ofNullable(this.method);
    }

    /**
     * The ID of an existing azure source.
     * 
     */
    @Import(name="resourceId")
    private @Nullable String resourceId;

    /**
     * @return The ID of an existing azure source.
     * 
     */
    public Optional<String> resourceId() {
        return Optional.ofNullable(this.resourceId);
    }

    /**
     * A list of path that needs to be exported from response body.
     * Setting it to `[&#34;*&#34;]` will export the full response body.
     * Here&#39;s an example. If it sets to `[&#34;keys&#34;]`, it will set the following json to computed property `output`.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *     }
     * }
     * ```
     * 
     */
    @Import(name="responseExportValues")
    private @Nullable List<String> responseExportValues;

    /**
     * @return A list of path that needs to be exported from response body.
     * Setting it to `[&#34;*&#34;]` will export the full response body.
     * Here&#39;s an example. If it sets to `[&#34;keys&#34;]`, it will set the following json to computed property `output`.
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *     }
     * }
     * ```
     * 
     */
    public Optional<List<String>> responseExportValues() {
        return Optional.ofNullable(this.responseExportValues);
    }

    /**
     * It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `&lt;api-version&gt;` is version of the API used to manage this azure resource.
     * 
     */
    @Import(name="type", required=true)
    private String type;

    /**
     * @return It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `&lt;api-version&gt;` is version of the API used to manage this azure resource.
     * 
     */
    public String type() {
        return this.type;
    }

    private GetResourceActionPlainArgs() {}

    private GetResourceActionPlainArgs(GetResourceActionPlainArgs $) {
        this.action = $.action;
        this.body = $.body;
        this.method = $.method;
        this.resourceId = $.resourceId;
        this.responseExportValues = $.responseExportValues;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetResourceActionPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetResourceActionPlainArgs $;

        public Builder() {
            $ = new GetResourceActionPlainArgs();
        }

        public Builder(GetResourceActionPlainArgs defaults) {
            $ = new GetResourceActionPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param action The name of the resource action. It&#39;s also possible to make Http requests towards the resource ID if leave this field empty.
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable String action) {
            $.action = action;
            return this;
        }

        /**
         * @param body A JSON object that contains the request body.
         * 
         * @return builder
         * 
         */
        public Builder body(@Nullable String body) {
            $.body = body;
            return this;
        }

        /**
         * @param method Specifies the Http method of the azure resource action. Allowed values are `POST` and `GET`. Defaults to `POST`.
         * 
         * @return builder
         * 
         */
        public Builder method(@Nullable String method) {
            $.method = method;
            return this;
        }

        /**
         * @param resourceId The ID of an existing azure source.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(@Nullable String resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param responseExportValues A list of path that needs to be exported from response body.
         * Setting it to `[&#34;*&#34;]` will export the full response body.
         * Here&#39;s an example. If it sets to `[&#34;keys&#34;]`, it will set the following json to computed property `output`.
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
         * import java.util.List;
         * import java.util.ArrayList;
         * import java.util.Map;
         * import java.io.File;
         * import java.nio.file.Files;
         * import java.nio.file.Paths;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder responseExportValues(@Nullable List<String> responseExportValues) {
            $.responseExportValues = responseExportValues;
            return this;
        }

        /**
         * @param responseExportValues A list of path that needs to be exported from response body.
         * Setting it to `[&#34;*&#34;]` will export the full response body.
         * Here&#39;s an example. If it sets to `[&#34;keys&#34;]`, it will set the following json to computed property `output`.
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
         * import java.util.List;
         * import java.util.ArrayList;
         * import java.util.Map;
         * import java.io.File;
         * import java.nio.file.Files;
         * import java.nio.file.Paths;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder responseExportValues(String... responseExportValues) {
            return responseExportValues(List.of(responseExportValues));
        }

        /**
         * @param type It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
         * `&lt;api-version&gt;` is version of the API used to manage this azure resource.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            $.type = type;
            return this;
        }

        public GetResourceActionPlainArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
