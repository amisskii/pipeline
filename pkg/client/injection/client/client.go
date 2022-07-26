/*
Copyright 2020 The Tekton Authors

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

// Code generated by injection-gen. DO NOT EDIT.

package client

import (
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"

	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	v1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	versioned "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	typedtektonv1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1"
	typedtektonv1alpha1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1alpha1"
	typedtektonv1beta1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	discovery "k8s.io/client-go/discovery"
	dynamic "k8s.io/client-go/dynamic"
	rest "k8s.io/client-go/rest"
	injection "knative.dev/pkg/injection"
	dynamicclient "knative.dev/pkg/injection/clients/dynamicclient"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterClient(withClientFromConfig)
	injection.Default.RegisterClientFetcher(func(ctx context.Context) interface{} {
		return Get(ctx)
	})
	injection.Dynamic.RegisterDynamicClient(withClientFromDynamic)
}

// Key is used as the key for associating information with a context.Context.
type Key struct{}

func withClientFromConfig(ctx context.Context, cfg *rest.Config) context.Context {
	return context.WithValue(ctx, Key{}, versioned.NewForConfigOrDie(cfg))
}

func withClientFromDynamic(ctx context.Context) context.Context {
	return context.WithValue(ctx, Key{}, &wrapClient{dyn: dynamicclient.Get(ctx)})
}

// Get extracts the versioned.Interface client from the context.
func Get(ctx context.Context) versioned.Interface {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		if injection.GetConfig(ctx) == nil {
			logging.FromContext(ctx).Panic(
				"Unable to fetch github.com/tektoncd/pipeline/pkg/client/clientset/versioned.Interface from context. This context is not the application context (which is typically given to constructors via sharedmain).")
		} else {
			logging.FromContext(ctx).Panic(
				"Unable to fetch github.com/tektoncd/pipeline/pkg/client/clientset/versioned.Interface from context.")
		}
	}
	return untyped.(versioned.Interface)
}

type wrapClient struct {
	dyn dynamic.Interface
}

var _ versioned.Interface = (*wrapClient)(nil)

func (w *wrapClient) Discovery() discovery.DiscoveryInterface {
	panic("Discovery called on dynamic client!")
}

func convert(from interface{}, to runtime.Object) error {
	bs, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("Marshal() = %w", err)
	}
	if err := json.Unmarshal(bs, to); err != nil {
		return fmt.Errorf("Unmarshal() = %w", err)
	}
	return nil
}

// TektonV1alpha1 retrieves the TektonV1alpha1Client
func (w *wrapClient) TektonV1alpha1() typedtektonv1alpha1.TektonV1alpha1Interface {
	return &wrapTektonV1alpha1{
		dyn: w.dyn,
	}
}

type wrapTektonV1alpha1 struct {
	dyn dynamic.Interface
}

func (w *wrapTektonV1alpha1) RESTClient() rest.Interface {
	panic("RESTClient called on dynamic client!")
}

func (w *wrapTektonV1alpha1) Runs(namespace string) typedtektonv1alpha1.RunInterface {
	return &wrapTektonV1alpha1RunImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "tekton.dev",
			Version:  "v1alpha1",
			Resource: "runs",
		}),

		namespace: namespace,
	}
}

type wrapTektonV1alpha1RunImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedtektonv1alpha1.RunInterface = (*wrapTektonV1alpha1RunImpl)(nil)

func (w *wrapTektonV1alpha1RunImpl) Create(ctx context.Context, in *v1alpha1.Run, opts v1.CreateOptions) (*v1alpha1.Run, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1alpha1",
		Kind:    "Run",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.Run{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1alpha1RunImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapTektonV1alpha1RunImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapTektonV1alpha1RunImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Run, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.Run{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1alpha1RunImpl) List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RunList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.RunList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1alpha1RunImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Run, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.Run{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1alpha1RunImpl) Update(ctx context.Context, in *v1alpha1.Run, opts v1.UpdateOptions) (*v1alpha1.Run, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1alpha1",
		Kind:    "Run",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.Run{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1alpha1RunImpl) UpdateStatus(ctx context.Context, in *v1alpha1.Run, opts v1.UpdateOptions) (*v1alpha1.Run, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1alpha1",
		Kind:    "Run",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1alpha1.Run{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1alpha1RunImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

// TektonV1beta1 retrieves the TektonV1beta1Client
func (w *wrapClient) TektonV1beta1() typedtektonv1beta1.TektonV1beta1Interface {
	return &wrapTektonV1beta1{
		dyn: w.dyn,
	}
}

type wrapTektonV1beta1 struct {
	dyn dynamic.Interface
}

func (w *wrapTektonV1beta1) RESTClient() rest.Interface {
	panic("RESTClient called on dynamic client!")
}

func (w *wrapTektonV1beta1) ClusterTasks() typedtektonv1beta1.ClusterTaskInterface {
	return &wrapTektonV1beta1ClusterTaskImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "tekton.dev",
			Version:  "v1beta1",
			Resource: "clustertasks",
		}),
	}
}

type wrapTektonV1beta1ClusterTaskImpl struct {
	dyn dynamic.NamespaceableResourceInterface
}

var _ typedtektonv1beta1.ClusterTaskInterface = (*wrapTektonV1beta1ClusterTaskImpl)(nil)

func (w *wrapTektonV1beta1ClusterTaskImpl) Create(ctx context.Context, in *v1beta1.ClusterTask, opts v1.CreateOptions) (*v1beta1.ClusterTask, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "ClusterTask",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.ClusterTask{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1ClusterTaskImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Delete(ctx, name, opts)
}

func (w *wrapTektonV1beta1ClusterTaskImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapTektonV1beta1ClusterTaskImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ClusterTask, error) {
	uo, err := w.dyn.Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.ClusterTask{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1ClusterTaskImpl) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ClusterTaskList, error) {
	uo, err := w.dyn.List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.ClusterTaskList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1ClusterTaskImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ClusterTask, err error) {
	uo, err := w.dyn.Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.ClusterTask{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1ClusterTaskImpl) Update(ctx context.Context, in *v1beta1.ClusterTask, opts v1.UpdateOptions) (*v1beta1.ClusterTask, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "ClusterTask",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.ClusterTask{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1ClusterTaskImpl) UpdateStatus(ctx context.Context, in *v1beta1.ClusterTask, opts v1.UpdateOptions) (*v1beta1.ClusterTask, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "ClusterTask",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.ClusterTask{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1ClusterTaskImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapTektonV1beta1) Pipelines(namespace string) typedtektonv1beta1.PipelineInterface {
	return &wrapTektonV1beta1PipelineImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "tekton.dev",
			Version:  "v1beta1",
			Resource: "pipelines",
		}),

		namespace: namespace,
	}
}

type wrapTektonV1beta1PipelineImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedtektonv1beta1.PipelineInterface = (*wrapTektonV1beta1PipelineImpl)(nil)

func (w *wrapTektonV1beta1PipelineImpl) Create(ctx context.Context, in *v1beta1.Pipeline, opts v1.CreateOptions) (*v1beta1.Pipeline, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "Pipeline",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.Pipeline{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapTektonV1beta1PipelineImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapTektonV1beta1PipelineImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Pipeline, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.Pipeline{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineImpl) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.PipelineList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.PipelineList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Pipeline, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.Pipeline{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineImpl) Update(ctx context.Context, in *v1beta1.Pipeline, opts v1.UpdateOptions) (*v1beta1.Pipeline, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "Pipeline",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.Pipeline{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineImpl) UpdateStatus(ctx context.Context, in *v1beta1.Pipeline, opts v1.UpdateOptions) (*v1beta1.Pipeline, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "Pipeline",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.Pipeline{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapTektonV1beta1) PipelineRuns(namespace string) typedtektonv1beta1.PipelineRunInterface {
	return &wrapTektonV1beta1PipelineRunImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "tekton.dev",
			Version:  "v1beta1",
			Resource: "pipelineruns",
		}),

		namespace: namespace,
	}
}

type wrapTektonV1beta1PipelineRunImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedtektonv1beta1.PipelineRunInterface = (*wrapTektonV1beta1PipelineRunImpl)(nil)

func (w *wrapTektonV1beta1PipelineRunImpl) Create(ctx context.Context, in *v1beta1.PipelineRun, opts v1.CreateOptions) (*v1beta1.PipelineRun, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "PipelineRun",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.PipelineRun{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineRunImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapTektonV1beta1PipelineRunImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapTektonV1beta1PipelineRunImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.PipelineRun, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.PipelineRun{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineRunImpl) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.PipelineRunList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.PipelineRunList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineRunImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PipelineRun, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.PipelineRun{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineRunImpl) Update(ctx context.Context, in *v1beta1.PipelineRun, opts v1.UpdateOptions) (*v1beta1.PipelineRun, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "PipelineRun",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.PipelineRun{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineRunImpl) UpdateStatus(ctx context.Context, in *v1beta1.PipelineRun, opts v1.UpdateOptions) (*v1beta1.PipelineRun, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "PipelineRun",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.PipelineRun{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1PipelineRunImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapTektonV1beta1) Tasks(namespace string) typedtektonv1beta1.TaskInterface {
	return &wrapTektonV1beta1TaskImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "tekton.dev",
			Version:  "v1beta1",
			Resource: "tasks",
		}),

		namespace: namespace,
	}
}

type wrapTektonV1beta1TaskImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedtektonv1beta1.TaskInterface = (*wrapTektonV1beta1TaskImpl)(nil)

func (w *wrapTektonV1beta1TaskImpl) Create(ctx context.Context, in *v1beta1.Task, opts v1.CreateOptions) (*v1beta1.Task, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "Task",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.Task{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapTektonV1beta1TaskImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapTektonV1beta1TaskImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Task, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.Task{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskImpl) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.TaskList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.TaskList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Task, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.Task{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskImpl) Update(ctx context.Context, in *v1beta1.Task, opts v1.UpdateOptions) (*v1beta1.Task, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "Task",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.Task{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskImpl) UpdateStatus(ctx context.Context, in *v1beta1.Task, opts v1.UpdateOptions) (*v1beta1.Task, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "Task",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.Task{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapTektonV1beta1) TaskRuns(namespace string) typedtektonv1beta1.TaskRunInterface {
	return &wrapTektonV1beta1TaskRunImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "tekton.dev",
			Version:  "v1beta1",
			Resource: "taskruns",
		}),

		namespace: namespace,
	}
}

type wrapTektonV1beta1TaskRunImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedtektonv1beta1.TaskRunInterface = (*wrapTektonV1beta1TaskRunImpl)(nil)

func (w *wrapTektonV1beta1TaskRunImpl) Create(ctx context.Context, in *v1beta1.TaskRun, opts v1.CreateOptions) (*v1beta1.TaskRun, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "TaskRun",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.TaskRun{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskRunImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapTektonV1beta1TaskRunImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapTektonV1beta1TaskRunImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.TaskRun, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.TaskRun{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskRunImpl) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.TaskRunList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.TaskRunList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskRunImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.TaskRun, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.TaskRun{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskRunImpl) Update(ctx context.Context, in *v1beta1.TaskRun, opts v1.UpdateOptions) (*v1beta1.TaskRun, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "TaskRun",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.TaskRun{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskRunImpl) UpdateStatus(ctx context.Context, in *v1beta1.TaskRun, opts v1.UpdateOptions) (*v1beta1.TaskRun, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1beta1",
		Kind:    "TaskRun",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &v1beta1.TaskRun{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1beta1TaskRunImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

// TektonV1 retrieves the TektonV1Client
func (w *wrapClient) TektonV1() typedtektonv1.TektonV1Interface {
	return &wrapTektonV1{
		dyn: w.dyn,
	}
}

type wrapTektonV1 struct {
	dyn dynamic.Interface
}

func (w *wrapTektonV1) RESTClient() rest.Interface {
	panic("RESTClient called on dynamic client!")
}

func (w *wrapTektonV1) Pipelines(namespace string) typedtektonv1.PipelineInterface {
	return &wrapTektonV1PipelineImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "tekton.dev",
			Version:  "v1",
			Resource: "pipelines",
		}),

		namespace: namespace,
	}
}

type wrapTektonV1PipelineImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedtektonv1.PipelineInterface = (*wrapTektonV1PipelineImpl)(nil)

func (w *wrapTektonV1PipelineImpl) Create(ctx context.Context, in *pipelinev1.Pipeline, opts v1.CreateOptions) (*pipelinev1.Pipeline, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1",
		Kind:    "Pipeline",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.Pipeline{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1PipelineImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapTektonV1PipelineImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapTektonV1PipelineImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*pipelinev1.Pipeline, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.Pipeline{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1PipelineImpl) List(ctx context.Context, opts v1.ListOptions) (*pipelinev1.PipelineList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.PipelineList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1PipelineImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *pipelinev1.Pipeline, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.Pipeline{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1PipelineImpl) Update(ctx context.Context, in *pipelinev1.Pipeline, opts v1.UpdateOptions) (*pipelinev1.Pipeline, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1",
		Kind:    "Pipeline",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.Pipeline{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1PipelineImpl) UpdateStatus(ctx context.Context, in *pipelinev1.Pipeline, opts v1.UpdateOptions) (*pipelinev1.Pipeline, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1",
		Kind:    "Pipeline",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.Pipeline{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1PipelineImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}

func (w *wrapTektonV1) Tasks(namespace string) typedtektonv1.TaskInterface {
	return &wrapTektonV1TaskImpl{
		dyn: w.dyn.Resource(schema.GroupVersionResource{
			Group:    "tekton.dev",
			Version:  "v1",
			Resource: "tasks",
		}),

		namespace: namespace,
	}
}

type wrapTektonV1TaskImpl struct {
	dyn dynamic.NamespaceableResourceInterface

	namespace string
}

var _ typedtektonv1.TaskInterface = (*wrapTektonV1TaskImpl)(nil)

func (w *wrapTektonV1TaskImpl) Create(ctx context.Context, in *pipelinev1.Task, opts v1.CreateOptions) (*pipelinev1.Task, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1",
		Kind:    "Task",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Create(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.Task{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1TaskImpl) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return w.dyn.Namespace(w.namespace).Delete(ctx, name, opts)
}

func (w *wrapTektonV1TaskImpl) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	return w.dyn.Namespace(w.namespace).DeleteCollection(ctx, opts, listOpts)
}

func (w *wrapTektonV1TaskImpl) Get(ctx context.Context, name string, opts v1.GetOptions) (*pipelinev1.Task, error) {
	uo, err := w.dyn.Namespace(w.namespace).Get(ctx, name, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.Task{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1TaskImpl) List(ctx context.Context, opts v1.ListOptions) (*pipelinev1.TaskList, error) {
	uo, err := w.dyn.Namespace(w.namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.TaskList{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1TaskImpl) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *pipelinev1.Task, err error) {
	uo, err := w.dyn.Namespace(w.namespace).Patch(ctx, name, pt, data, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.Task{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1TaskImpl) Update(ctx context.Context, in *pipelinev1.Task, opts v1.UpdateOptions) (*pipelinev1.Task, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1",
		Kind:    "Task",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).Update(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.Task{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1TaskImpl) UpdateStatus(ctx context.Context, in *pipelinev1.Task, opts v1.UpdateOptions) (*pipelinev1.Task, error) {
	in.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "tekton.dev",
		Version: "v1",
		Kind:    "Task",
	})
	uo := &unstructured.Unstructured{}
	if err := convert(in, uo); err != nil {
		return nil, err
	}
	uo, err := w.dyn.Namespace(w.namespace).UpdateStatus(ctx, uo, opts)
	if err != nil {
		return nil, err
	}
	out := &pipelinev1.Task{}
	if err := convert(uo, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (w *wrapTektonV1TaskImpl) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return nil, errors.New("NYI: Watch")
}
