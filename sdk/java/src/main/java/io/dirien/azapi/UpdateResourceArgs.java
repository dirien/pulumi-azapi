// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.azapi;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UpdateResourceArgs extends com.pulumi.resources.ResourceArgs {

    public static final UpdateResourceArgs Empty = new UpdateResourceArgs();

    /**
     * A JSON object that contains the request body used to add on an existing azure resource.
     * 
     */
    @Import(name="body")
    private @Nullable Output<String> body;

    /**
     * @return A JSON object that contains the request body used to add on an existing azure resource.
     * 
     */
    public Optional<Output<String>> body() {
        return Optional.ofNullable(this.body);
    }

    /**
     * Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
     * 
     */
    @Import(name="ignoreCasing")
    private @Nullable Output<Boolean> ignoreCasing;

    /**
     * @return Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> ignoreCasing() {
        return Optional.ofNullable(this.ignoreCasing);
    }

    /**
     * Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
     * 
     */
    @Import(name="ignoreMissingProperty")
    private @Nullable Output<Boolean> ignoreMissingProperty;

    /**
     * @return Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> ignoreMissingProperty() {
        return Optional.ofNullable(this.ignoreMissingProperty);
    }

    /**
     * A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
     * 
     */
    @Import(name="locks")
    private @Nullable Output<List<String>> locks;

    /**
     * @return A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
     * 
     */
    public Optional<Output<List<String>>> locks() {
        return Optional.ofNullable(this.locks);
    }

    /**
     * Specifies the name of the azure resource. Changing this forces a new resource to be created.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Specifies the name of the azure resource. Changing this forces a new resource to be created.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="parentId")
    private @Nullable Output<String> parentId;

    public Optional<Output<String>> parentId() {
        return Optional.ofNullable(this.parentId);
    }

    /**
     * The ID of an existing azure source. Changing this forces a new azure resource to be created.
     * 
     */
    @Import(name="resourceId")
    private @Nullable Output<String> resourceId;

    /**
     * @return The ID of an existing azure source. Changing this forces a new azure resource to be created.
     * 
     */
    public Optional<Output<String>> resourceId() {
        return Optional.ofNullable(this.resourceId);
    }

    /**
     * A list of path that needs to be exported from response body.
     * Setting it to `[&#34;*&#34;]` will export the full response body.
     * Here&#39;s an example. If it sets to `[&#34;properties.loginServer&#34;, &#34;properties.policies.quarantinePolicy.status&#34;]`, it will set the following json to computed property `output`.
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
    private @Nullable Output<List<String>> responseExportValues;

    /**
     * @return A list of path that needs to be exported from response body.
     * Setting it to `[&#34;*&#34;]` will export the full response body.
     * Here&#39;s an example. If it sets to `[&#34;properties.loginServer&#34;, &#34;properties.policies.quarantinePolicy.status&#34;]`, it will set the following json to computed property `output`.
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
    public Optional<Output<List<String>>> responseExportValues() {
        return Optional.ofNullable(this.responseExportValues);
    }

    /**
     * It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `&lt;api-version&gt;` is version of the API used to manage this azure resource.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `&lt;api-version&gt;` is version of the API used to manage this azure resource.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private UpdateResourceArgs() {}

    private UpdateResourceArgs(UpdateResourceArgs $) {
        this.body = $.body;
        this.ignoreCasing = $.ignoreCasing;
        this.ignoreMissingProperty = $.ignoreMissingProperty;
        this.locks = $.locks;
        this.name = $.name;
        this.parentId = $.parentId;
        this.resourceId = $.resourceId;
        this.responseExportValues = $.responseExportValues;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UpdateResourceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UpdateResourceArgs $;

        public Builder() {
            $ = new UpdateResourceArgs();
        }

        public Builder(UpdateResourceArgs defaults) {
            $ = new UpdateResourceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param body A JSON object that contains the request body used to add on an existing azure resource.
         * 
         * @return builder
         * 
         */
        public Builder body(@Nullable Output<String> body) {
            $.body = body;
            return this;
        }

        /**
         * @param body A JSON object that contains the request body used to add on an existing azure resource.
         * 
         * @return builder
         * 
         */
        public Builder body(String body) {
            return body(Output.of(body));
        }

        /**
         * @param ignoreCasing Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ignoreCasing(@Nullable Output<Boolean> ignoreCasing) {
            $.ignoreCasing = ignoreCasing;
            return this;
        }

        /**
         * @param ignoreCasing Whether ignore incorrect casing returned in `body` to suppress plan-diff. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder ignoreCasing(Boolean ignoreCasing) {
            return ignoreCasing(Output.of(ignoreCasing));
        }

        /**
         * @param ignoreMissingProperty Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder ignoreMissingProperty(@Nullable Output<Boolean> ignoreMissingProperty) {
            $.ignoreMissingProperty = ignoreMissingProperty;
            return this;
        }

        /**
         * @param ignoreMissingProperty Whether ignore not returned properties like credentials in `body` to suppress plan-diff. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder ignoreMissingProperty(Boolean ignoreMissingProperty) {
            return ignoreMissingProperty(Output.of(ignoreMissingProperty));
        }

        /**
         * @param locks A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
         * 
         * @return builder
         * 
         */
        public Builder locks(@Nullable Output<List<String>> locks) {
            $.locks = locks;
            return this;
        }

        /**
         * @param locks A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
         * 
         * @return builder
         * 
         */
        public Builder locks(List<String> locks) {
            return locks(Output.of(locks));
        }

        /**
         * @param locks A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.
         * 
         * @return builder
         * 
         */
        public Builder locks(String... locks) {
            return locks(List.of(locks));
        }

        /**
         * @param name Specifies the name of the azure resource. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Specifies the name of the azure resource. Changing this forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder parentId(@Nullable Output<String> parentId) {
            $.parentId = parentId;
            return this;
        }

        public Builder parentId(String parentId) {
            return parentId(Output.of(parentId));
        }

        /**
         * @param resourceId The ID of an existing azure source. Changing this forces a new azure resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(@Nullable Output<String> resourceId) {
            $.resourceId = resourceId;
            return this;
        }

        /**
         * @param resourceId The ID of an existing azure source. Changing this forces a new azure resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder resourceId(String resourceId) {
            return resourceId(Output.of(resourceId));
        }

        /**
         * @param responseExportValues A list of path that needs to be exported from response body.
         * Setting it to `[&#34;*&#34;]` will export the full response body.
         * Here&#39;s an example. If it sets to `[&#34;properties.loginServer&#34;, &#34;properties.policies.quarantinePolicy.status&#34;]`, it will set the following json to computed property `output`.
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
        public Builder responseExportValues(@Nullable Output<List<String>> responseExportValues) {
            $.responseExportValues = responseExportValues;
            return this;
        }

        /**
         * @param responseExportValues A list of path that needs to be exported from response body.
         * Setting it to `[&#34;*&#34;]` will export the full response body.
         * Here&#39;s an example. If it sets to `[&#34;properties.loginServer&#34;, &#34;properties.policies.quarantinePolicy.status&#34;]`, it will set the following json to computed property `output`.
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
        public Builder responseExportValues(List<String> responseExportValues) {
            return responseExportValues(Output.of(responseExportValues));
        }

        /**
         * @param responseExportValues A list of path that needs to be exported from response body.
         * Setting it to `[&#34;*&#34;]` will export the full response body.
         * Here&#39;s an example. If it sets to `[&#34;properties.loginServer&#34;, &#34;properties.policies.quarantinePolicy.status&#34;]`, it will set the following json to computed property `output`.
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
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
         * `&lt;api-version&gt;` is version of the API used to manage this azure resource.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public UpdateResourceArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
