package vclusterconfig

import "testing"

func TestConvertPlatformConfig(t *testing.T) {
	type args struct {
		legacyPlatformConfig *LegacyPlatformConfig
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "empty config",
			args:    args{legacyPlatformConfig: &LegacyPlatformConfig{}},
			want:    "{}\n",
			wantErr: false,
		},
		{
			name: "AutoSnapshot basic",
			args: args{
				legacyPlatformConfig: &LegacyPlatformConfig{
					AutoSnapshot: &LegacyAutoSnapshot{
						Enabled:  true,
						Schedule: "0 */12 * * *",
					},
				},
			},
			want: `snapshots:
  auto:
    schedule: 0 */12 * * *
`,
			wantErr: false,
		},
		{
			name: "AutoDelete conversion",
			args: args{
				legacyPlatformConfig: &LegacyPlatformConfig{
					AutoDelete: &LegacyAutoDelete{AfterInactivity: 3600},
				},
			},
			want: `deletion:
  auto:
    afterInactivity: 1h0m0s
`,
			wantErr: false,
		},
		{
			name: "AutoSleep basic",
			args: args{
				legacyPlatformConfig: &LegacyPlatformConfig{
					AutoSleep: &LegacyPlatformAutoSleep{
						AfterInactivity: 1800,
						Schedule:        "0 22 * * *",
						Timezone:        "America/New_York",
					},
				},
			},
			want: `sleep:
  auto:
    afterInactivity: 30m0s
    schedule: 0 22 * * *
    timezone: America/New_York
`,
			wantErr: false,
		},
		{
			name: "AutoSleep with AutoWakeup",
			args: args{
				legacyPlatformConfig: &LegacyPlatformConfig{
					AutoSleep: &LegacyPlatformAutoSleep{
						AfterInactivity: 1800,
						AutoWakeup:      &LegacyPlatformAutoWakeup{Schedule: "0 8 * * *"},
					},
				},
			},
			want: `sleep:
  auto:
    afterInactivity: 30m0s
    wakeup:
      schedule: 0 8 * * *
`,
			wantErr: false,
		},
		{
			name: "Platform with Project only",
			args: args{
				legacyPlatformConfig: &LegacyPlatformConfig{
					Project: "my-project",
				},
			},
			want: `platform:
  project: my-project
`,
			wantErr: false,
		},
		{
			name: "Platform with APIKey",
			args: args{
				legacyPlatformConfig: &LegacyPlatformConfig{
					APIKey: map[string]any{
						"secretName": "my-secret",
						"namespace":  "my-namespace",
						"createRBAC": true,
					},
				},
			},
			want: `platform:
  apiKey:
    createRBAC: true
    namespace: my-namespace
    secretName: my-secret
`,
			wantErr: false,
		},
		{
			name: "Snapshots with S3 storage and retention",
			args: args{
				legacyPlatformConfig: &LegacyPlatformConfig{
					AutoSnapshot: &LegacyAutoSnapshot{
						Enabled:  true,
						Timezone: "UTC",
						Schedule: "0 0 * * *",
						Storage: LegacyScheduledSnapshotStorage{
							Type: "s3",
							S3: LegacyS3Storage{
								Url: "s3://my-bucket/snapshots",
								Credential: &SnapshotSecretCredential{
									SecretName:      "s3-creds",
									SecretNamespace: "default",
								},
							},
						},
						Retention: SnapshotRetention{Period: 30, MaxSnapshots: 10},
					},
				},
			},
			want: `snapshots:
  auto:
    retention:
      maxSnapshots: 10
      period: 30
    schedule: 0 0 * * *
    storage:
      s3:
        credential:
          secretName: s3-creds
          secretNamespace: default
        url: s3://my-bucket/snapshots
      type: s3
    timezone: UTC
`,
			wantErr: false,
		},
		{
			name: "Combined conversions",
			args: args{
				legacyPlatformConfig: &LegacyPlatformConfig{
					AutoDelete: &LegacyAutoDelete{AfterInactivity: 7200},
					Project:    "combined-project",
					AutoSnapshot: &LegacyAutoSnapshot{
						Enabled:  true,
						Schedule: "0 */6 * * *",
					},
				},
			},
			want: `deletion:
  auto:
    afterInactivity: 2h0m0s
platform:
  project: combined-project
snapshots:
  auto:
    schedule: 0 */6 * * *
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertPlatformConfig(tt.args.legacyPlatformConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertPlatformConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertPlatformConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
