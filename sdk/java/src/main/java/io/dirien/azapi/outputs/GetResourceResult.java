// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.azapi.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.azapi.outputs.GetResourceIdentity;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetResourceResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return An `identity` block as defined below, which contains the Managed Service Identity information for this azure resource.
     * 
     */
    private GetResourceIdentity identity;
    /**
     * @return The Azure Region where the azure resource should exist.
     * 
     */
    private String location;
    private @Nullable String name;
    /**
     * @return The output json containing the properties specified in `response_export_values`. Here&#39;re some examples to decode json and extract the value.
     * 
     */
    private String output;
    private @Nullable String parentId;
    private @Nullable String resourceId;
    private @Nullable List<String> responseExportValues;
    /**
     * @return A mapping of tags which should be assigned to the azure resource.
     * 
     */
    private Map<String,String> tags;
    /**
     * @return The Type of Identity which should be used for this azure resource. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned,UserAssigned`.
     * 
     */
    private String type;

    private GetResourceResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return An `identity` block as defined below, which contains the Managed Service Identity information for this azure resource.
     * 
     */
    public GetResourceIdentity identity() {
        return this.identity;
    }
    /**
     * @return The Azure Region where the azure resource should exist.
     * 
     */
    public String location() {
        return this.location;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return The output json containing the properties specified in `response_export_values`. Here&#39;re some examples to decode json and extract the value.
     * 
     */
    public String output() {
        return this.output;
    }
    public Optional<String> parentId() {
        return Optional.ofNullable(this.parentId);
    }
    public Optional<String> resourceId() {
        return Optional.ofNullable(this.resourceId);
    }
    public List<String> responseExportValues() {
        return this.responseExportValues == null ? List.of() : this.responseExportValues;
    }
    /**
     * @return A mapping of tags which should be assigned to the azure resource.
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }
    /**
     * @return The Type of Identity which should be used for this azure resource. Possible values are `SystemAssigned`, `UserAssigned` and `SystemAssigned,UserAssigned`.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetResourceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private GetResourceIdentity identity;
        private String location;
        private @Nullable String name;
        private String output;
        private @Nullable String parentId;
        private @Nullable String resourceId;
        private @Nullable List<String> responseExportValues;
        private Map<String,String> tags;
        private String type;
        public Builder() {}
        public Builder(GetResourceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.identity = defaults.identity;
    	      this.location = defaults.location;
    	      this.name = defaults.name;
    	      this.output = defaults.output;
    	      this.parentId = defaults.parentId;
    	      this.resourceId = defaults.resourceId;
    	      this.responseExportValues = defaults.responseExportValues;
    	      this.tags = defaults.tags;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder identity(GetResourceIdentity identity) {
            this.identity = Objects.requireNonNull(identity);
            return this;
        }
        @CustomType.Setter
        public Builder location(String location) {
            this.location = Objects.requireNonNull(location);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder output(String output) {
            this.output = Objects.requireNonNull(output);
            return this;
        }
        @CustomType.Setter
        public Builder parentId(@Nullable String parentId) {
            this.parentId = parentId;
            return this;
        }
        @CustomType.Setter
        public Builder resourceId(@Nullable String resourceId) {
            this.resourceId = resourceId;
            return this;
        }
        @CustomType.Setter
        public Builder responseExportValues(@Nullable List<String> responseExportValues) {
            this.responseExportValues = responseExportValues;
            return this;
        }
        public Builder responseExportValues(String... responseExportValues) {
            return responseExportValues(List.of(responseExportValues));
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetResourceResult build() {
            final var o = new GetResourceResult();
            o.id = id;
            o.identity = identity;
            o.location = location;
            o.name = name;
            o.output = output;
            o.parentId = parentId;
            o.resourceId = resourceId;
            o.responseExportValues = responseExportValues;
            o.tags = tags;
            o.type = type;
            return o;
        }
    }
}