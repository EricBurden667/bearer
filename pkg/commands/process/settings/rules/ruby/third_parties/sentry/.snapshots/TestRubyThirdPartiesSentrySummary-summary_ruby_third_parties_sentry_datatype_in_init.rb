critical:
    - policy_name: ""
      policy_dsrid: DSR-1
      policy_display_id: ruby_third_parties_sentry
      policy_description: Do not send sensitive data to Sentry.
      line_number: 3
      filename: pkg/commands/process/settings/rules/ruby/third_parties/sentry/testdata/datatype_in_init.rb
      category_groups:
        - PII
      parent_line_number: 3
      parent_content: 'breadcrumb.message = "Authenticated user #{current_user.email}"'


--
