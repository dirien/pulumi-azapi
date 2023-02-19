// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.azapi.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.azapi.inputs.ResourceIdentityArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResourceState extends com.pulumi.resources.ResourceArgs {

    public static final ResourceState Empty = new ResourceState();

    /**
     * A JSON object that contains the request body used to create and update azure resource.
     * 
     */
    @Import(name="body")
    private @Nullable Output<String> body;

    /**
     * @return A JSON object that contains the request body used to create and update azure resource.
     * 
     */
    public Optional<Output<String>> body() {
        return Optional.ofNullable(this.body);
    }

    /**
     * A `identity` block as defined below.
     * 
     */
    @Import(name="identity")
    private @Nullable Output<ResourceIdentityArgs> identity;

    /**
     * @return A `identity` block as defined below.
     * 
     */
    public Optional<Output<ResourceIdentityArgs>> identity() {
        return Optional.ofNullable(this.identity);
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
     * The Azure Region where the azure resource should exist.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The Azure Region where the azure resource should exist.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
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

    /**
     * The output json containing the properties specified in `response_export_values`. Here&#39;re some examples to decode json and extract the value.
     * 
     */
    @Import(name="output")
    private @Nullable Output<String> output;

    /**
     * @return The output json containing the properties specified in `response_export_values`. Here&#39;re some examples to decode json and extract the value.
     * 
     */
    public Optional<Output<String>> output() {
        return Optional.ofNullable(this.output);
    }

    @Import(name="parentId")
    private @Nullable Output<String> parentId;

    public Optional<Output<String>> parentId() {
        return Optional.ofNullable(this.parentId);
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
     * Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
     * 
     */
    @Import(name="schemaValidationEnabled")
    private @Nullable Output<Boolean> schemaValidationEnabled;

    /**
     * @return Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> schemaValidationEnabled() {
        return Optional.ofNullable(this.schemaValidationEnabled);
    }

    /**
     * A mapping of tags which should be assigned to the azure resource.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A mapping of tags which should be assigned to the azure resource.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `&lt;api-version&gt;` is version of the API used to manage this azure resource.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
     * `&lt;api-version&gt;` is version of the API used to manage this azure resource.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private ResourceState() {}

    private ResourceState(ResourceState $) {
        this.body = $.body;
        this.identity = $.identity;
        this.ignoreCasing = $.ignoreCasing;
        this.ignoreMissingProperty = $.ignoreMissingProperty;
        this.location = $.location;
        this.locks = $.locks;
        this.name = $.name;
        this.output = $.output;
        this.parentId = $.parentId;
        this.responseExportValues = $.responseExportValues;
        this.schemaValidationEnabled = $.schemaValidationEnabled;
        this.tags = $.tags;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResourceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResourceState $;

        public Builder() {
            $ = new ResourceState();
        }

        public Builder(ResourceState defaults) {
            $ = new ResourceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param body A JSON object that contains the request body used to create and update azure resource.
         * 
         * @return builder
         * 
         */
        public Builder body(@Nullable Output<String> body) {
            $.body = body;
            return this;
        }

        /**
         * @param body A JSON object that contains the request body used to create and update azure resource.
         * 
         * @return builder
         * 
         */
        public Builder body(String body) {
            return body(Output.of(body));
        }

        /**
         * @param identity A `identity` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder identity(@Nullable Output<ResourceIdentityArgs> identity) {
            $.identity = identity;
            return this;
        }

        /**
         * @param identity A `identity` block as defined below.
         * 
         * @return builder
         * 
         */
        public Builder identity(ResourceIdentityArgs identity) {
            return identity(Output.of(identity));
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
         * @param location The Azure Region where the azure resource should exist.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The Azure Region where the azure resource should exist.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
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

        /**
         * @param output The output json containing the properties specified in `response_export_values`. Here&#39;re some examples to decode json and extract the value.
         * 
         * @return builder
         * 
         */
        public Builder output(@Nullable Output<String> output) {
            $.output = output;
            return this;
        }

        /**
         * @param output The output json containing the properties specified in `response_export_values`. Here&#39;re some examples to decode json and extract the value.
         * 
         * @return builder
         * 
         */
        public Builder output(String output) {
            return output(Output.of(output));
        }

        public Builder parentId(@Nullable Output<String> parentId) {
            $.parentId = parentId;
            return this;
        }

        public Builder parentId(String parentId) {
            return parentId(Output.of(parentId));
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
         * @param schemaValidationEnabled Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder schemaValidationEnabled(@Nullable Output<Boolean> schemaValidationEnabled) {
            $.schemaValidationEnabled = schemaValidationEnabled;
            return this;
        }

        /**
         * @param schemaValidationEnabled Whether enabled the validation on `type` and `body` with embedded schema. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder schemaValidationEnabled(Boolean schemaValidationEnabled) {
            return schemaValidationEnabled(Output.of(schemaValidationEnabled));
        }

        /**
         * @param tags A mapping of tags which should be assigned to the azure resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A mapping of tags which should be assigned to the azure resource.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param type It is in a format like `&lt;resource-type&gt;@&lt;api-version&gt;`. `&lt;resource-type&gt;` is the Azure resource type, for example, `Microsoft.Storage/storageAccounts`.
         * `&lt;api-version&gt;` is version of the API used to manage this azure resource.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
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

        public ResourceState build() {
            return $;
        }
    }

}