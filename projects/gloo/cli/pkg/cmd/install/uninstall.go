package install

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/pkg/cliutil/install"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
)

func UninstallGloo(opts *options.Options, cli install.KubeCli) error {
	if err := uninstallGloo(opts, cli); err != nil {
		fmt.Fprintf(os.Stderr, "Uninstall failed. Detailed logs available at %s.\n", cliutil.GetLogsPath())
		return err
	}
	return nil
}

func uninstallGloo(opts *options.Options, cli install.KubeCli) error {
	if opts.Uninstall.DeleteNamespace || opts.Uninstall.DeleteAll {
		if err := deleteNamespace(cli, opts.Uninstall.Namespace); err != nil {
			return err
		}
	} else {
		if err := deleteGlooSystem(cli, opts.Uninstall.Namespace); err != nil {
			return err
		}
	}

	if opts.Uninstall.DeleteCrds || opts.Uninstall.DeleteAll {
		if err := deleteGlooCrds(cli); err != nil {
			return err
		}
	}

	if opts.Uninstall.DeleteAll {
		if err := deleteRbac(cli); err != nil {
			return err
		}
	}

	if err := uninstallKnativeIfNecessary(); err != nil {
		return err
	}
	return nil
}

func deleteRbac(cli install.KubeCli) error {
	fmt.Printf("Removing Gloo RBAC configuration...\n")
	for _, rbacKind := range GlooRbacKinds {
		if err := cli.Kubectl(nil, "delete", rbacKind, "-l", "app=gloo"); err != nil {
			return errors.Wrapf(err, "deleting rbac failed")
		}
	}
	return nil
}

func deleteGlooSystem(cli install.KubeCli, namespace string) error {
	fmt.Printf("Removing Gloo system components from namespace %s...\n", namespace)
	for _, kind := range GlooSystemKinds {
		if err := cli.Kubectl(nil, "delete", kind, "-l", "app=gloo", "-n", namespace); err != nil {
			return errors.Wrapf(err, "deleting gloo system failed")
		}
	}
	return nil
}

func deleteGlooCrds(cli install.KubeCli) error {
	fmt.Printf("Removing Gloo CRDs...\n")
	args := []string{"delete", "crd"}
	for _, crd := range GlooCrdNames {
		args = append(args, crd)
	}
	if err := cli.Kubectl(nil, args...); err != nil {
		return errors.Wrapf(err, "deleting crds failed")
	}
	return nil
}

func deleteNamespace(cli install.KubeCli, namespace string) error {
	fmt.Printf("Removing namespace %s...\n", namespace)
	if err := cli.Kubectl(nil, "delete", "namespace", namespace); err != nil {
		return errors.Wrapf(err, "deleting namespace %s failed", namespace)
	}
	return nil
}

func uninstallKnativeIfNecessary() error {
	_, installOpts, err := checkKnativeInstallation()
	if err != nil {
		return errors.Wrapf(err, "finding knative installation")
	}
	if installOpts != nil {
		fmt.Printf("Removing knative components installed by Gloo %#v...\n", installOpts)
		manifests, err := RenderKnativeManifests(*installOpts)
		if err != nil {
			return errors.Wrapf(err, "rendering knative manifests")
		}
		if err := install.KubectlDelete([]byte(manifests), "--ignore-not-found"); err != nil {
			return errors.Wrapf(err, "deleting knative failed")
		}
	}
	return nil
}
