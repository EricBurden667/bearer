high:
    - rule_dsrid: DSR-1
      rule_display_id: javascript_third_parties_airbrake
      rule_description: Do not send sensitive data to Airbrake.
      rule_documentation_url: https://curio.sh/reference/rules/javascript_third_parties_airbrake
      line_number: 18
      filename: pkg/commands/process/settings/rules/javascript/third_parties/airbrake/testdata/datatype_in_notify.js
      category_groups:
        - Personal Data
      parent_line_number: 16
      parent_content: |-
        airbrake.notify({
              error: err,
              params: { user: user.ipAddress },
            })


--
