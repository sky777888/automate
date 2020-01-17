##### GRPC SETUP #####
require 'api/reporting/reporting_pb'
require 'api/reporting/reporting_services_pb'

describe File.basename(__FILE__) do
  Reporting = Chef::Automate::Domain::Compliance::Api::Reporting unless defined?(Reporting)

  def reporting
    Reporting::ReportingService;
  end

  it "no filters gets them all filter" do
    actual_data = GRPC reporting, :list_control_items, Reporting::ControlItemRequest.new(
        filters: [
            Reporting::ListFilter.new(type: 'end_time', values: ['2018-03-04T23:59:59Z'])
        ]
    )
    control_item_array = actual_data['control_items']
    assert_equal(18, control_item_array.size)
  end

  it "suggests control tag keys matching 'scope'" do
    actual_data = GRPC reporting, :list_control_items, Reporting::ControlItemRequest.new(
        filters: [
            Reporting::ListFilter.new(type: 'start_time', values: ['2018-02-01T23:59:59Z']),
            Reporting::ListFilter.new(type: 'end_time', values: ['2018-03-04T23:59:59Z'])
        ],
        size: 10
    )
    control_item_array = actual_data['control_items']
    assert_equal('apache-01', control_item_array[0]['id'])
    assert_equal('DevSec Apache Baseline', control_item_array[0]['profile']['title'])
    assert_equal(10, control_item_array.size)
  end

  it "control list items with a size of 2" do
    actual_data = GRPC reporting, :list_control_items, Reporting::ControlItemRequest.new(
        filters: [
            Reporting::ListFilter.new(type: 'start_time', values: ['2018-02-01T23:59:59Z']),
            Reporting::ListFilter.new(type: 'end_time', values: ['2018-03-04T23:59:59Z'])
        ],
        size: 2
    )

    expected_data = {
        "controlItems" => [
            {
                "id" => "apache-01",
                "title" => "Apache should be running",
                "profile" => {
                    "title" => "DevSec Apache Baseline",
                    "id" => "41a02784bfea15592ba2748d55927d8d1f9da205816ef18d3bb2ebe4c5ce18a9",
                    "version" => "2.0.1"
                }, "impact" => 1,
                "endTime" => "2018-03-04T09:18:42Z",
                "controlSummary" => {
                    "total" => 5,
                    "passed" => {},
                    "skipped" => {
                        "total" => 5
                    },
                    "failed" => {}
                }
            },
            {
                "id" => "apache-02",
                "title" => "Apache should be enabled",
                "profile" => {
                    "title" => "DevSec Apache Baseline",
                    "id" => "41a02784bfea15592ba2748d55927d8d1f9da205816ef18d3bb2ebe4c5ce18a9",
                    "version" => "2.0.1"
                },
                "impact" => 1,
                "endTime" => "2018-03-04T09:18:42Z",
                "controlSummary" => {
                    "total" => 5,
                    "passed" => {},
                    "skipped" => {
                        "total" => 4
                    },
                    "failed" => {
                        "total" => 1,
                        "critical" => 1
                    }
                }
            }
        ]
    }
    control_item_array = actual_data['control_items']
    assert_equal(2, control_item_array.size)
    assert_equal_json_content(expected_data, actual_data)
  end


  it "control list items with a size of 2 with nginx profile" do
    actual_data = GRPC reporting, :list_control_items, Reporting::ControlItemRequest.new(
        filters: [
            Reporting::ListFilter.new(type: 'start_time', values: ['2018-02-01T23:59:59Z']),
            Reporting::ListFilter.new(type: "profile_id", values: ["09adcbb3b9b3233d5de63cd98a5ba3e155b3aaeb66b5abed379f5fb1ff143988"]),
            Reporting::ListFilter.new(type: 'end_time', values: ['2018-03-04T23:59:59Z']),
            Reporting::ListFilter.new(type: "platform", values: ["centos"])

        ],
        size: 2
    )
    expected_data = {
        "controlItems" => [
            {
                "id" => "nginx-01",
                "title" => "Running worker process as non-privileged user",
                "profile" => {
                    "title" => "DevSec Nginx Baseline",
                    "id" => "09adcbb3b9b3233d5de63cd98a5ba3e155b3aaeb66b5abed379f5fb1ff143988",
                    "version" => "2.1.0"
                },
                "impact" => 1,
                "endTime" => "2018-03-04T09:18:41Z",
                "controlSummary" => {
                    "total" => 1,
                    "passed" => {
                        "total" => 1
                    },
                    "skipped" => {},
                    "failed" => {}
                }
            },
            {
                "id" => "nginx-02",
                "title" => "Check NGINX config file owner, group and permissions.",
                "profile" => {
                    "title" => "DevSec Nginx Baseline",
                    "id" => "09adcbb3b9b3233d5de63cd98a5ba3e155b3aaeb66b5abed379f5fb1ff143988",
                    "version" => "2.1.0"
                },
                "impact" => 1,
                "endTime" => "2018-03-04T09:18:41Z",
                "controlSummary" => {
                    "total" => 1,
                    "passed" => {
                        "total" => 1
                    },
                    "skipped" => {},
                    "failed" => {}
                }
            }
        ]
    }

    control_item_array = actual_data['control_items']
    assert_equal(2, control_item_array.size)
    assert_equal_json_content(expected_data, actual_data)
  end

  it "control list items with a size of 2 but only care about skipped" do
    actual_data = GRPC reporting, :list_control_items, Reporting::ControlItemRequest.new(
        filters: [
            Reporting::ListFilter.new(type: 'start_time', values: ['2018-02-01T23:59:59Z']),
            Reporting::ListFilter.new(type: 'end_time', values: ['2018-03-04T23:59:59Z']),
            Reporting::ListFilter.new(type: 'control_status', values: ['skipped']),
        ],
        size: 2
    )

    expected_data = {
        "controlItems" => [
            {
                "id" => "apache-01",
                "title" => "Apache should be running",
                "profile" => {
                    "title" => "DevSec Apache Baseline",
                    "id" => "41a02784bfea15592ba2748d55927d8d1f9da205816ef18d3bb2ebe4c5ce18a9",
                    "version" => "2.0.1"
                }, "impact" => 1,
                "endTime" => "2018-03-04T09:18:42Z",
                "controlSummary" => {
                    "total" => 5,
                    "passed" => {},
                    "skipped" => {
                        "total" => 5
                    },
                    "failed" => {}
                }
            },
            {
                "id" => "apache-02",
                "title" => "Apache should be enabled",
                "profile" => {
                    "title" => "DevSec Apache Baseline",
                    "id" => "41a02784bfea15592ba2748d55927d8d1f9da205816ef18d3bb2ebe4c5ce18a9",
                    "version" => "2.0.1"
                },
                "impact" => 1,
                "endTime" => "2018-03-04T09:18:41Z",
                "controlSummary" => {
                    "total" => 4,
                    "passed" => {},
                    "skipped" => {
                        "total" => 4
                    },
                    "failed" => {}
                }
            }
        ]
    }

    control_item_array = actual_data['control_items']
    assert_equal(2, control_item_array.size)
    assert_equal_json_content(expected_data, actual_data)
  end

  it "control list items with a size of 2 but only care about failed" do
    actual_data = GRPC reporting, :list_control_items, Reporting::ControlItemRequest.new(
        filters: [
            Reporting::ListFilter.new(type: 'start_time', values: ['2018-02-01T23:59:59Z']),
            Reporting::ListFilter.new(type: 'end_time', values: ['2018-03-04T23:59:59Z']),
            Reporting::ListFilter.new(type: 'control_status', values: ['failed']),
        ],
        size: 2
    )

    expected_data = {
        "controlItems" => [
            {
                "id" => "apache-02",
                "title" => "Apache should be enabled",
                "profile" => {
                    "title" => "DevSec Apache Baseline",
                    "id" => "41a02784bfea15592ba2748d55927d8d1f9da205816ef18d3bb2ebe4c5ce18a9",
                    "version" => "2.0.1"
                },
                "impact" => 1,
                "endTime" => "2018-03-04T09:18:42Z",
                "controlSummary" => {
                    "total" => 1,
                    "passed" => {},
                    "skipped" => {},
                    "failed" => {
                        "total" => 1,
                        "critical" => 1
                    }
                }
            },
            {
                "id" => "nginx-01",
                "title" => "Running worker process as non-privileged user",
                "profile" => {
                    "title" => "DevSec Nginx Baseline",
                    "id" => "09adcbb3b9b3233d5de63cd98a5ba3e155b3aaeb66b5abed379f5fb1ff143988",
                    "version" => "2.1.0"
                },
                "impact" => 1,
                "endTime" => "2018-03-04T09:18:42Z",
                "controlSummary" => {
                    "total" => 2,
                    "passed" => {},
                    "skipped" => {},
                    "failed" => {
                        "total" => 2,
                        "critical" => 2
                    }
                }
            }
        ]
    }

    control_item_array = actual_data['control_items']
    assert_equal(2, control_item_array.size)
    assert_equal_json_content(expected_data, actual_data)
  end

  it "control list items with a size of 2 but only care about failed" do
    actual_data = GRPC reporting, :list_control_items, Reporting::ControlItemRequest.new(
        filters: [
            Reporting::ListFilter.new(type: 'start_time', values: ['2018-02-01T23:59:59Z']),
            Reporting::ListFilter.new(type: 'end_time', values: ['2018-03-04T23:59:59Z']),
            Reporting::ListFilter.new(type: 'control_tag:satisfies', values: ['NGX-1', 'SRG-00006'])
        ],
        size: 2
    )

    expected_data = {
        "controlItems" => [
            {
                "id" => "nginx-01",
                "title" => "Running worker process as non-privileged user",
                "profile" => {
                    "title" => "DevSec Nginx Baseline",
                    "id" => "09adcbb3b9b3233d5de63cd98a5ba3e155b3aaeb66b5abed379f5fb1ff143988",
                    "version" => "2.1.0"
                },
                "impact" => 1,
                "endTime" => "2018-03-04T09:18:43Z",
                "controlSummary" => {
                    "total" => 1,
                    "passed" => {},
                    "skipped" => {},
                    "failed" => {
                        "total" => 1,
                        "critical" => 1
                    }
                }
            }
        ]
    }
    control_item_array = actual_data['control_items']
    assert_equal(1, control_item_array.size)
    assert_equal_json_content(expected_data, actual_data)
  end

  it "control list items with a size of 4 filtered by control_tag" do
    actual_data = GRPC reporting, :list_control_items, Reporting::ControlItemRequest.new(
        filters: [
            Reporting::ListFilter.new(type: 'start_time', values: ['2018-03-01T23:59:59Z']),
            Reporting::ListFilter.new(type: 'end_time', values: ['2018-03-04T23:59:59Z']),
            Reporting::ListFilter.new(type: 'control_tag:scoop', values: ['icecream'])
        ],
        size: 4
    )
    expected_data = {
        "controlItems" =>
            [
                {
                    "id" => "apache-01",
                    "title" => "Apache should be running",
                    "profile" => {
                        "title" => "DevSec Apache Baseline",
                        "id" => "41a02784bfea15592ba2748d55927d8d1f9da205816ef18d3bb2ebe4c5ce18a9",
                        "version" => "2.0.1"
                    },
                    "impact" => 1,
                    "endTime" => "2018-03-04T09:18:41Z",
                    "controlSummary" => {
                        "total" => 1,
                        "passed" => {},
                        "skipped" => {
                            "total" => 1
                        },
                        "failed" => {}
                    }
                }
            ]
    }
    assert_equal_json_content(expected_data, actual_data)
  end

  it "control list items filtered by environment and platform" do
    actual_data = GRPC reporting, :list_control_items, Reporting::ControlItemRequest.new(
        filters: [
            Reporting::ListFilter.new(type: 'start_time', values: ['2018-03-01T23:59:59Z']),
            Reporting::ListFilter.new(type: 'end_time', values: ['2018-03-04T23:59:59Z']),
            Reporting::ListFilter.new(type: 'environment', values: ['DevSec Prod Zeta', 'missing']),
            Reporting::ListFilter.new(type: 'platform', values: ['windows', 'missing'])
        ],
        size: 1000
    )
    control_item_array = actual_data['control_items']
    assert_equal(14, control_item_array.size)
  end

  it "control list items with a size of 2" do
    actual_data = GRPC reporting, :list_control_items, Reporting::ControlItemRequest.new(
        filters: [
            Reporting::ListFilter.new(type: 'start_time', values: ['2018-04-01T00:00:00Z']),
            Reporting::ListFilter.new(type: 'end_time', values: ['2018-04-01T23:59:59Z'])
        ],
        size: 10
    )

    expected_data = {
      "controlItems" => [
        {
          "id" => "apache-01",
          "title" => "Apache should be running",
          "profile" => {
              "title" => "DevSec Apache Baseline",
              "id" => "41a02784bfea15592ba2748d55927d8d1f9da205816ef18d3bb2ebe4c5ce18a9",
              "version" => "2.0.1"
          }, "impact" => 1,
          "endTime" => "2018-03-04T09:18:42Z",
          "waived_nodes" => [
            {
              "waived_str" => "no",
              "controlSummary" => { "total" => 8, "passed" => {}, "skipped" => {}, "failed" => { "total" => 8 }}
            },
            {
              "waived_str" => "no_expired",
              "controlSummary" => { "total" => 7, "passed" => {}, "skipped" => {}, "failed" => { "total" => 7 }}
            },
            {
              "waived_str" => "yes_run",
              "expiration_date" => "2022-02-02",
              "justification" => "Fix until 2022",
              "controlSummary" => { "total" => 6, "passed" => { "total" => 6 }, "skipped" => {}, "failed" => {}}
            },
            {
              "waived_str" => "yes_run",
              "expiration_date" => "2021-01-01",
              "justification" => "Fix until 2021",
              "controlSummary" => { "total" => 5, "passed" => {}, "skipped" => {}, "failed" => { "total" => 5 }}
            },
            {
              "waived_str" => "yes",
              "expiration_date" => "2025-01-01",
              "justification" => "Sound reasoning 25",
              "controlSummary" => { "total" => 4, "passed" => {}, "skipped" => { "total" => 4 }, "failed" => {}}
            },
            {
              "waived_str" => "yes",
              "expiration_date" => "2026-01-01",
              "justification" => "Sound reasoning 26",
              "controlSummary" => { "total" => 3, "passed" => {}, "skipped" => { "total" => 3 }, "failed" => {}}
            }
          ]
        }
      ]
    }
    control_item_array = actual_data['control_items']
    assert_equal(2, control_item_array.size)
    assert_equal_json_content(expected_data, actual_data)
  end

end
