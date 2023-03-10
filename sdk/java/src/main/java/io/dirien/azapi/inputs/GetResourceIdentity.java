// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.azapi.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetResourceIdentity extends com.pulumi.resources.InvokeArgs {

    public static final GetResourceIdentity Empty = new GetResourceIdentity();

    /**
     * A list of User Managed Identity ID&#39;s which should be assigned to the azure resource.
     * 
     */
    @Import(name="identityIds", required=true)
    private List<String> identityIds;

    /**
     * @return A list of User Managed Identity ID&#39;s which should be assigned to the azure resource.
     * 
     */
    public List<String> identityIds() {
        return this.identityIds;
    }

    /**
     * The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
     * 
     */
    @Import(name="principalId", required=true)
    private String principalId;

    /**
     * @return The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
     * 
     */
    public String principalId() {
        return this.principalId;
    }

    /**
     * The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
     * 
     */
    @Import(name="tenantId", required=true)
    private String tenantId;

    /**
     * @return The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
     * 
     */
    public String tenantId() {
        return this.tenantId;
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

    private GetResourceIdentity() {}

    private GetResourceIdentity(GetResourceIdentity $) {
        this.identityIds = $.identityIds;
        this.principalId = $.principalId;
        this.tenantId = $.tenantId;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetResourceIdentity defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetResourceIdentity $;

        public Builder() {
            $ = new GetResourceIdentity();
        }

        public Builder(GetResourceIdentity defaults) {
            $ = new GetResourceIdentity(Objects.requireNonNull(defaults));
        }

        /**
         * @param identityIds A list of User Managed Identity ID&#39;s which should be assigned to the azure resource.
         * 
         * @return builder
         * 
         */
        public Builder identityIds(List<String> identityIds) {
            $.identityIds = identityIds;
            return this;
        }

        /**
         * @param identityIds A list of User Managed Identity ID&#39;s which should be assigned to the azure resource.
         * 
         * @return builder
         * 
         */
        public Builder identityIds(String... identityIds) {
            return identityIds(List.of(identityIds));
        }

        /**
         * @param principalId The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.
         * 
         * @return builder
         * 
         */
        public Builder principalId(String principalId) {
            $.principalId = principalId;
            return this;
        }

        /**
         * @param tenantId The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource.
         * 
         * @return builder
         * 
         */
        public Builder tenantId(String tenantId) {
            $.tenantId = tenantId;
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
            $.type = type;
            return this;
        }

        public GetResourceIdentity build() {
            $.identityIds = Objects.requireNonNull($.identityIds, "expected parameter 'identityIds' to be non-null");
            $.principalId = Objects.requireNonNull($.principalId, "expected parameter 'principalId' to be non-null");
            $.tenantId = Objects.requireNonNull($.tenantId, "expected parameter 'tenantId' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
