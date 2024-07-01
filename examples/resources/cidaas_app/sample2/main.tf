module "shared_settings" {
  source = "./modules"
  common_config = {
    client_type                       = "SINGLE_PAGE"
    accent_color                      = "#ef4923"                        // Default: #ef4923
    primary_color                     = "#ef4923"                        // Default: #f7941d
    media_type                        = "IMAGE"                          // Default: IMAGE
    allow_login_with                  = ["EMAIL", "MOBILE", "USER_NAME"] // Default: ["EMAIL", "MOBILE", "USER_NAME"]
    redirect_uris                     = ["https://cidaas.com"]
    allowed_logout_urls               = ["https://cidaas.com"]
    enable_deduplication              = true      // Default: false
    auto_login_after_register         = true      // Default: false
    enable_passwordless_auth          = false     // Default: true
    register_with_login_information   = false     // Default: false
    fds_enabled                       = false     // Default: true
    hosted_page_group                 = "default" // Default: default
    company_name                      = "Widas ID GmbH"
    company_address                   = "01"
    company_website                   = "https://cidaas.com"
    allowed_scopes                    = ["openid", "cidaas:register", "profile"]
    response_types                    = ["code", "token", "id_token"] // Default: ["code", "token", "id_token"]
    grant_types                       = ["client_credentials"]        // Default: ["implicit", "authorization_code", "password", "refresh_token"]
    login_providers                   = ["login_provider1", "login_provider2"]
    is_hybrid_app                     = true // Default: false
    allowed_web_origins               = ["https://cidaas.com"]
    allowed_origins                   = ["https://cidaas.com"]
    default_max_age                   = 86400      // Default: 86400
    token_lifetime_in_seconds         = 86400      // Default: 86400
    id_token_lifetime_in_seconds      = 86400      // Default: 86400
    refresh_token_lifetime_in_seconds = 15780000   // Default: 15780000
    template_group_id                 = "custtemp" // Default: default
    editable                          = true       // Default: true
    social_providers = [{
      provider_name = "cidaas social provider"
      social_id     = "fdc63bd0-6044-4fa0-abff"
      display_name  = "cidaas"
    }]
    custom_providers = [{
      logo_url      = "https://cidaas.com/logo-url"
      provider_name = "sample-custom-provider"
      display_name  = "sample-custom-provider"
      type          = "CUSTOM_OPENID_CONNECT"
    }]
    saml_providers = [{
      logo_url      = "https://cidaas.com/logo-url"
      provider_name = "sample-sampl-provider"
      display_name  = "sample-sampl-provider"
      type          = "SAMPL_IDP_PROVIDER"
    }]
    ad_providers = [{
      logo_url      = "https://cidaas.com/logo-url"
      provider_name = "sample-ad-provider"
      display_name  = "sample-ad-provider"
      type          = "ADD_PROVIDER"
    }]
    jwe_enabled  = true // Default: false
    user_consent = true // Default: false
    allowed_groups = [{
      group_id      = "developer101"
      roles         = ["developer", "qa", "admin"]
      default_roles = ["developer"]
    }]
    operations_allowed_groups = [{
      group_id      = "developer101"
      roles         = ["developer", "qa", "admin"]
      default_roles = ["developer"]
    }]
    enabled                     = true // Default: true
    always_ask_mfa              = true // Default: false
    allowed_mfa                 = ["OFF"]
    email_verification_required = true // Default: true
    allowed_roles               = ["sample"]
    default_roles               = ["sample"]
    enable_classical_provider   = true     // Default: true
    is_remember_me_selected     = true     // Default: true
    bot_provider                = "CIDAAS" // Default: CIDAAS
    allow_guest_login           = true     // Default: false
    #  mfa Default:
    #  {
    #   setting = "OFF"
    #  }
    mfa = {
      setting = "OFF"
    }
    webfinger      = "no_redirection"
    default_scopes = ["sample"]
    pending_scopes = ["sample"]
  }
}
