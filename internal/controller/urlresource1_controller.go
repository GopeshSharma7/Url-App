/*
Copyright 2023.

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

package controller

import (
	"context"
	// "crypto/sha256"
	"encoding/base64"
	"fmt"
	demov1 "project1/api/v1"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// UrlResource1Reconciler reconciles a UrlResource1 object
type UrlResource1Reconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.demo.kcd.io,resources=urlresource1s,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.demo.kcd.io,resources=urlresource1s/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.demo.kcd.io,resources=urlresource1s/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the UrlResource1 object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *UrlResource1Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("urlshortener", req.NamespacedName)

	// Fetch the URLShortener custom resource.
	var urlShortener demov1.UrlResource1
	if err := r.Get(ctx, req.NamespacedName, &urlShortener); err != nil {
		if errors.IsNotFound(err) {
			log.Info("URLShortener resource not found, it may have been deleted, and it will get new")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get URLShortener")
		return ctrl.Result{}, err
	}

	// Reconciliation logic: Generate a "shortURL" based on the targetURL.
	// In a real-world scenario, you would implement proper URL shortening here.
	//shortURL := generateShortURL(urlShortener.Spec.TargetURL)
	urlMap := make(map[string]string)
	for _, targetURL := range urlShortener.Spec.TargetURLs {
		shortURL := generateShortURL(targetURL)
		urlMap[targetURL] = shortURL
	}

	// Update the URLMap field in the Status to store the short URLs.
	urlShortener.Status.URLMap = urlMap
	//fmt.Println("received encoded value:", shortURL)
	// Update the status with the generated shortURL.
	//urlShortener.Status.ShortURL = shortURLs
	if err := r.Status().Update(ctx, &urlShortener); err != nil {
		log.Error(err, "Failed to update URLShortener status")
		return ctrl.Result{}, err
	}

	fmt.Println("Reconciliation completed", "ShortURL", urlMap)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *UrlResource1Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.UrlResource1{}).
		Complete(r)
}

func generateShortURL(targetURL string) string {
	// h := sha256.New()
	// h.Write([]byte(targetURL))
	// hash := h.Sum(nil)
	// return fmt.Sprintf("%x", hash[len(hash)-8:])

	encodedURL := base64.StdEncoding.EncodeToString([]byte(targetURL))
	//fmt.Println("========#######encoded value : ", encodedURL[len(encodedURL)-8:])
	return fmt.Sprintf("%v", encodedURL[len(encodedURL)-8:])
}
