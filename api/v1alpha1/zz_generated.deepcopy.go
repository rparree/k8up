// +build !ignore_autogenerated

/*
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/operator-framework/operator-lib/status"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Archive) DeepCopyInto(out *Archive) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Archive.
func (in *Archive) DeepCopy() *Archive {
	if in == nil {
		return nil
	}
	out := new(Archive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Archive) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveList) DeepCopyInto(out *ArchiveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Archive, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveList.
func (in *ArchiveList) DeepCopy() *ArchiveList {
	if in == nil {
		return nil
	}
	out := new(ArchiveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ArchiveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveSchedule) DeepCopyInto(out *ArchiveSchedule) {
	*out = *in
	in.ArchiveSpec.DeepCopyInto(&out.ArchiveSpec)
	if in.ScheduleCommon != nil {
		in, out := &in.ScheduleCommon, &out.ScheduleCommon
		*out = new(ScheduleCommon)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveSchedule.
func (in *ArchiveSchedule) DeepCopy() *ArchiveSchedule {
	if in == nil {
		return nil
	}
	out := new(ArchiveSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArchiveSpec) DeepCopyInto(out *ArchiveSpec) {
	*out = *in
	if in.RestoreSpec != nil {
		in, out := &in.RestoreSpec, &out.RestoreSpec
		*out = new(RestoreSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArchiveSpec.
func (in *ArchiveSpec) DeepCopy() *ArchiveSpec {
	if in == nil {
		return nil
	}
	out := new(ArchiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSpec) DeepCopyInto(out *AzureSpec) {
	*out = *in
	if in.AccountNameSecretRef != nil {
		in, out := &in.AccountNameSecretRef, &out.AccountNameSecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountKeySecretRef != nil {
		in, out := &in.AccountKeySecretRef, &out.AccountKeySecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSpec.
func (in *AzureSpec) DeepCopy() *AzureSpec {
	if in == nil {
		return nil
	}
	out := new(AzureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *B2Spec) DeepCopyInto(out *B2Spec) {
	*out = *in
	if in.AccountIDSecretRef != nil {
		in, out := &in.AccountIDSecretRef, &out.AccountIDSecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AccountKeySecretRef != nil {
		in, out := &in.AccountKeySecretRef, &out.AccountKeySecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new B2Spec.
func (in *B2Spec) DeepCopy() *B2Spec {
	if in == nil {
		return nil
	}
	out := new(B2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	if in.RepoPasswordSecretRef != nil {
		in, out := &in.RepoPasswordSecretRef, &out.RepoPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalSpec)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		*out = new(GCSSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Swift != nil {
		in, out := &in.Swift, &out.Swift
		*out = new(SwiftSpec)
		**out = **in
	}
	if in.B2 != nil {
		in, out := &in.B2, &out.B2
		*out = new(B2Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Rest != nil {
		in, out := &in.Rest, &out.Rest
		*out = new(RestServerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backup) DeepCopyInto(out *Backup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backup.
func (in *Backup) DeepCopy() *Backup {
	if in == nil {
		return nil
	}
	out := new(Backup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupList) DeepCopyInto(out *BackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupList.
func (in *BackupList) DeepCopy() *BackupList {
	if in == nil {
		return nil
	}
	out := new(BackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSchedule) DeepCopyInto(out *BackupSchedule) {
	*out = *in
	in.BackupSpec.DeepCopyInto(&out.BackupSpec)
	if in.ScheduleCommon != nil {
		in, out := &in.ScheduleCommon, &out.ScheduleCommon
		*out = new(ScheduleCommon)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSchedule.
func (in *BackupSchedule) DeepCopy() *BackupSchedule {
	if in == nil {
		return nil
	}
	out := new(BackupSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	in.RunnableSpec.DeepCopyInto(&out.RunnableSpec)
	if in.KeepJobs != nil {
		in, out := &in.KeepJobs, &out.KeepJobs
		*out = new(int)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupTemplate) DeepCopyInto(out *BackupTemplate) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	in.Backend.DeepCopyInto(&out.Backend)
	out.Env = in.Env
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupTemplate.
func (in *BackupTemplate) DeepCopy() *BackupTemplate {
	if in == nil {
		return nil
	}
	out := new(BackupTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Check) DeepCopyInto(out *Check) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Check.
func (in *Check) DeepCopy() *Check {
	if in == nil {
		return nil
	}
	out := new(Check)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Check) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckList) DeepCopyInto(out *CheckList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Check, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckList.
func (in *CheckList) DeepCopy() *CheckList {
	if in == nil {
		return nil
	}
	out := new(CheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CheckList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckSchedule) DeepCopyInto(out *CheckSchedule) {
	*out = *in
	in.CheckSpec.DeepCopyInto(&out.CheckSpec)
	if in.ScheduleCommon != nil {
		in, out := &in.ScheduleCommon, &out.ScheduleCommon
		*out = new(ScheduleCommon)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckSchedule.
func (in *CheckSchedule) DeepCopy() *CheckSchedule {
	if in == nil {
		return nil
	}
	out := new(CheckSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckSpec) DeepCopyInto(out *CheckSpec) {
	*out = *in
	in.RunnableSpec.DeepCopyInto(&out.RunnableSpec)
	if in.KeepJobs != nil {
		in, out := &in.KeepJobs, &out.KeepJobs
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckSpec.
func (in *CheckSpec) DeepCopy() *CheckSpec {
	if in == nil {
		return nil
	}
	out := new(CheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Env) DeepCopyInto(out *Env) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Env.
func (in *Env) DeepCopy() *Env {
	if in == nil {
		return nil
	}
	out := new(Env)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FolderRestore) DeepCopyInto(out *FolderRestore) {
	*out = *in
	if in.PersistentVolumeClaimVolumeSource != nil {
		in, out := &in.PersistentVolumeClaimVolumeSource, &out.PersistentVolumeClaimVolumeSource
		*out = new(v1.PersistentVolumeClaimVolumeSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FolderRestore.
func (in *FolderRestore) DeepCopy() *FolderRestore {
	if in == nil {
		return nil
	}
	out := new(FolderRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCSSpec) DeepCopyInto(out *GCSSpec) {
	*out = *in
	if in.ProjectIDSecretRef != nil {
		in, out := &in.ProjectIDSecretRef, &out.ProjectIDSecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.AccessTokenSecretRef != nil {
		in, out := &in.AccessTokenSecretRef, &out.AccessTokenSecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCSSpec.
func (in *GCSSpec) DeepCopy() *GCSSpec {
	if in == nil {
		return nil
	}
	out := new(GCSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSpec) DeepCopyInto(out *LocalSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSpec.
func (in *LocalSpec) DeepCopy() *LocalSpec {
	if in == nil {
		return nil
	}
	out := new(LocalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pod) DeepCopyInto(out *Pod) {
	*out = *in
	in.PodTemplateSpec.DeepCopyInto(&out.PodTemplateSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pod.
func (in *Pod) DeepCopy() *Pod {
	if in == nil {
		return nil
	}
	out := new(Pod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreBackupPod) DeepCopyInto(out *PreBackupPod) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreBackupPod.
func (in *PreBackupPod) DeepCopy() *PreBackupPod {
	if in == nil {
		return nil
	}
	out := new(PreBackupPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreBackupPod) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreBackupPodList) DeepCopyInto(out *PreBackupPodList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PreBackupPod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreBackupPodList.
func (in *PreBackupPodList) DeepCopy() *PreBackupPodList {
	if in == nil {
		return nil
	}
	out := new(PreBackupPodList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreBackupPodList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreBackupPodSpec) DeepCopyInto(out *PreBackupPodSpec) {
	*out = *in
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(Pod)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreBackupPodSpec.
func (in *PreBackupPodSpec) DeepCopy() *PreBackupPodSpec {
	if in == nil {
		return nil
	}
	out := new(PreBackupPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prune) DeepCopyInto(out *Prune) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prune.
func (in *Prune) DeepCopy() *Prune {
	if in == nil {
		return nil
	}
	out := new(Prune)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Prune) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PruneList) DeepCopyInto(out *PruneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Prune, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PruneList.
func (in *PruneList) DeepCopy() *PruneList {
	if in == nil {
		return nil
	}
	out := new(PruneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PruneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PruneSchedule) DeepCopyInto(out *PruneSchedule) {
	*out = *in
	in.PruneSpec.DeepCopyInto(&out.PruneSpec)
	if in.ScheduleCommon != nil {
		in, out := &in.ScheduleCommon, &out.ScheduleCommon
		*out = new(ScheduleCommon)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PruneSchedule.
func (in *PruneSchedule) DeepCopy() *PruneSchedule {
	if in == nil {
		return nil
	}
	out := new(PruneSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PruneSpec) DeepCopyInto(out *PruneSpec) {
	*out = *in
	in.RunnableSpec.DeepCopyInto(&out.RunnableSpec)
	in.Retention.DeepCopyInto(&out.Retention)
	if in.KeepJobs != nil {
		in, out := &in.KeepJobs, &out.KeepJobs
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PruneSpec.
func (in *PruneSpec) DeepCopy() *PruneSpec {
	if in == nil {
		return nil
	}
	out := new(PruneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestServerSpec) DeepCopyInto(out *RestServerSpec) {
	*out = *in
	if in.UserSecretRef != nil {
		in, out := &in.UserSecretRef, &out.UserSecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PasswordSecretReg != nil {
		in, out := &in.PasswordSecretReg, &out.PasswordSecretReg
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestServerSpec.
func (in *RestServerSpec) DeepCopy() *RestServerSpec {
	if in == nil {
		return nil
	}
	out := new(RestServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Restore) DeepCopyInto(out *Restore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Restore.
func (in *Restore) DeepCopy() *Restore {
	if in == nil {
		return nil
	}
	out := new(Restore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Restore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreList) DeepCopyInto(out *RestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Restore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreList.
func (in *RestoreList) DeepCopy() *RestoreList {
	if in == nil {
		return nil
	}
	out := new(RestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreMethod) DeepCopyInto(out *RestoreMethod) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3Spec)
		(*in).DeepCopyInto(*out)
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(FolderRestore)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreMethod.
func (in *RestoreMethod) DeepCopy() *RestoreMethod {
	if in == nil {
		return nil
	}
	out := new(RestoreMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreSchedule) DeepCopyInto(out *RestoreSchedule) {
	*out = *in
	in.RestoreSpec.DeepCopyInto(&out.RestoreSpec)
	if in.ScheduleCommon != nil {
		in, out := &in.ScheduleCommon, &out.ScheduleCommon
		*out = new(ScheduleCommon)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreSchedule.
func (in *RestoreSchedule) DeepCopy() *RestoreSchedule {
	if in == nil {
		return nil
	}
	out := new(RestoreSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestoreSpec) DeepCopyInto(out *RestoreSpec) {
	*out = *in
	in.RunnableSpec.DeepCopyInto(&out.RunnableSpec)
	if in.RestoreMethod != nil {
		in, out := &in.RestoreMethod, &out.RestoreMethod
		*out = new(RestoreMethod)
		(*in).DeepCopyInto(*out)
	}
	if in.KeepJobs != nil {
		in, out := &in.KeepJobs, &out.KeepJobs
		*out = new(int)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestoreSpec.
func (in *RestoreSpec) DeepCopy() *RestoreSpec {
	if in == nil {
		return nil
	}
	out := new(RestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicy) DeepCopyInto(out *RetentionPolicy) {
	*out = *in
	if in.KeepTags != nil {
		in, out := &in.KeepTags, &out.KeepTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicy.
func (in *RetentionPolicy) DeepCopy() *RetentionPolicy {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnableSpec) DeepCopyInto(out *RunnableSpec) {
	*out = *in
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(Backend)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnableSpec.
func (in *RunnableSpec) DeepCopy() *RunnableSpec {
	if in == nil {
		return nil
	}
	out := new(RunnableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Spec) DeepCopyInto(out *S3Spec) {
	*out = *in
	if in.AccessKeyIDSecretRef != nil {
		in, out := &in.AccessKeyIDSecretRef, &out.AccessKeyIDSecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretAccessKeySecretRef != nil {
		in, out := &in.SecretAccessKeySecretRef, &out.SecretAccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Spec.
func (in *S3Spec) DeepCopy() *S3Spec {
	if in == nil {
		return nil
	}
	out := new(S3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schedule) DeepCopyInto(out *Schedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schedule.
func (in *Schedule) DeepCopy() *Schedule {
	if in == nil {
		return nil
	}
	out := new(Schedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Schedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleCommon) DeepCopyInto(out *ScheduleCommon) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleCommon.
func (in *ScheduleCommon) DeepCopy() *ScheduleCommon {
	if in == nil {
		return nil
	}
	out := new(ScheduleCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleList) DeepCopyInto(out *ScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Schedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleList.
func (in *ScheduleList) DeepCopy() *ScheduleList {
	if in == nil {
		return nil
	}
	out := new(ScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleSpec) DeepCopyInto(out *ScheduleSpec) {
	*out = *in
	if in.Restore != nil {
		in, out := &in.Restore, &out.Restore
		*out = new(RestoreSchedule)
		(*in).DeepCopyInto(*out)
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(BackupSchedule)
		(*in).DeepCopyInto(*out)
	}
	if in.Archive != nil {
		in, out := &in.Archive, &out.Archive
		*out = new(ArchiveSchedule)
		(*in).DeepCopyInto(*out)
	}
	if in.Check != nil {
		in, out := &in.Check, &out.Check
		*out = new(CheckSchedule)
		(*in).DeepCopyInto(*out)
	}
	if in.Prune != nil {
		in, out := &in.Prune, &out.Prune
		*out = new(PruneSchedule)
		(*in).DeepCopyInto(*out)
	}
	if in.Backend != nil {
		in, out := &in.Backend, &out.Backend
		*out = new(Backend)
		(*in).DeepCopyInto(*out)
	}
	if in.KeepJobs != nil {
		in, out := &in.KeepJobs, &out.KeepJobs
		*out = new(int)
		**out = **in
	}
	in.ResourceRequirementsTemplate.DeepCopyInto(&out.ResourceRequirementsTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleSpec.
func (in *ScheduleSpec) DeepCopy() *ScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleStatus) DeepCopyInto(out *ScheduleStatus) {
	*out = *in
	if in.EffectiveSchedules != nil {
		in, out := &in.EffectiveSchedules, &out.EffectiveSchedules
		*out = make(map[JobType]ScheduleDefinition, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleStatus.
func (in *ScheduleStatus) DeepCopy() *ScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Snapshot) DeepCopyInto(out *Snapshot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Snapshot.
func (in *Snapshot) DeepCopy() *Snapshot {
	if in == nil {
		return nil
	}
	out := new(Snapshot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Snapshot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotList) DeepCopyInto(out *SnapshotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Snapshot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotList.
func (in *SnapshotList) DeepCopy() *SnapshotList {
	if in == nil {
		return nil
	}
	out := new(SnapshotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSpec) DeepCopyInto(out *SnapshotSpec) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Date != nil {
		in, out := &in.Date, &out.Date
		*out = (*in).DeepCopy()
	}
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSpec.
func (in *SnapshotSpec) DeepCopy() *SnapshotSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotStatus) DeepCopyInto(out *SnapshotStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotStatus.
func (in *SnapshotStatus) DeepCopy() *SnapshotStatus {
	if in == nil {
		return nil
	}
	out := new(SnapshotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftSpec) DeepCopyInto(out *SwiftSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftSpec.
func (in *SwiftSpec) DeepCopy() *SwiftSpec {
	if in == nil {
		return nil
	}
	out := new(SwiftSpec)
	in.DeepCopyInto(out)
	return out
}
