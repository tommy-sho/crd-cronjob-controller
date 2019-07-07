/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//並列で起動する際のポリシー
type ConcurrencyPolicy string

const (
	// AllowConcurrent allows CronJobs to run concurrently.
	AllowConcurrent ConcurrencyPolicy = "Allow"

	// ForbidConcurrent forbids concurrent runs, skipping next run if previous
	// hasn't finished yet.
	ForbidConcurrent ConcurrencyPolicy = "Forbid"

	// ReplaceConcurrent cancels currently running job and replaces it with a new one.
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)


// CronJobSpec defines the desired state of CronJob
type CronJobSpec struct {
	// Cronフォーマットのスケジュール
	Schedule string `json:"schedule"`

	// Jobがスケジュール通りに起動できなかったときのDeadline。optional
	// +optional
	StartingDeadlineSeconds *int64 `json:"startingDeadlineSeconds,omitempty"`

	// Jobが並列に起動された場合の振る舞いを定義する
	// - "Allow" (default): 並列で起動することを許可;
	// - "Forbid": 前のJobが起動していた場合、スキップする;
	// - "Replace":現在起動しているJobをキャンセルして新しいJobを起動させる
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrencyPolicy,omitempty"`

	// 後続Jobを実行するかどうかのフラグ、Defaultではfalse
	// +optional
	Suspend *bool `json:"suspend,omitempty"`

	// CronJobが起動するJobのテンプレート.
	JobTemplate batchv1beta1.JobTemplateSpec `json:"jobTemplate"`

	// 正常に完了したJobの数を保持する。明示的な0と未指定を区別するためにポインタになっている
	// +optional
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit,omitempty"`

	// 正常に完了しなかったJobの数を保持する。以下同文
	// +optional
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit,omitempty"`
}


// CronJobStatus defines the observed state of CronJob
type CronJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// 起動中のJobのステータス
	// +optional
	Active []corev1.ObjectReference `json:"active,omitempty"`

	//最後にスケジュールに成功した時間。time.timeではなくmeta/v1.time型である点に注意
	// +optional
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`
}


// +kubebuilder:object:root=true

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
