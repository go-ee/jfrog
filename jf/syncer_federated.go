package jf

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
)

func BuildFederatedRepoCloner(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager) (ret RepositoryCloner, err error) {

	switch packageType {
	case Bower:
		ret = NewBowerFederatedRepositoryCloner(source, target)
	case Chef:
		ret = NewChefFederatedRepositoryCloner(source, target)
	case CocoaPods:
		ret = NewCocoaPodsFederatedPodsRepositoryCloner(source, target)
	case Conan:
		ret = NewConanFederatedRepositoryCloner(source, target)
	case Docker:
		ret = NewDockerFederatedRepositoryCloner(source, target)
	case Go:
		ret = NewGoFederatedRepositoryCloner(source, target)
	case NuGet:
		ret = NewNuGetFederatedRepositoryCloner(source, target)
	case Npm:
		ret = NewNpmFederatedRepositoryCloner(source, target)
	case PhpComposer:
		ret = NewPhpComposerFederatedRepositoryCloner(source, target)
	case Puppet:
		ret = NewPuppetFederatedRepositoryCloner(source, target)
	case PyPi:
		ret = NewPyPiFederatedRepositoryCloner(source, target)
	case RubyGems:
		ret = NewRubyFederatedGemsRepositoryCloner(source, target)
	case Generic:
		ret = NewGenericFederatedRepositoryCloner(source, target)
	case Maven:
		ret = NewMavenFederatedRepositoryCloner(source, target)
	case Helm:
		ret = NewHelmFederatedRepositoryCloner(source, target)
	case GitLfs:
		ret = NewGitLfsFederatedRepositoryCloner(source, target)
	case Debian:
		ret = NewDebianFederatedRepositoryCloner(source, target)
	case YUM:
		ret = NewYumFederatedRepositoryCloner(source, target)
	case Vagrant:
		ret = NewVagrantFederatedRepositoryCloner(source, target)
	case Cargo:
		ret = NewCargoFederatedRepositoryCloner(source, target)
	case Gradle:
		ret = NewGradleFederatedRepositoryCloner(source, target)
	default:
		err = fmt.Errorf("repo type '%v' with package type '%v' not supported", Federated, packageType)

	}
	return
}

func NewFederatedRepositoryClonerImpl(packageType PackageType,
	source *ArtifactoryManager, target *ArtifactoryManager) *RepositoryClonerImpl {
	return NewRepositoryCloner(Federated, packageType, source, target)
}

type BowerFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewBowerFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *BowerFederatedRepositoryCloner {
	return &BowerFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Bower, source, target)}
}

func (o *BowerFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.BowerFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Bower(sourceRepoDetails)
		})
	}
	return
}

type ChefFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewChefFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *ChefFederatedRepositoryCloner {
	return &ChefFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Chef, source, target)}
}

func (o *ChefFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ChefFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Chef(sourceRepoDetails)
		})
	}
	return
}

type CocoaPodsFederatedPodsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCocoaPodsFederatedPodsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *CocoaPodsFederatedPodsRepositoryCloner {
	return &CocoaPodsFederatedPodsRepositoryCloner{
		NewFederatedRepositoryClonerImpl(CocoaPods, source, target)}
}

func (o *CocoaPodsFederatedPodsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CocoapodsFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Cocoapods(sourceRepoDetails)
		})
	}
	return
}

type ConanFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewConanFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *ConanFederatedRepositoryCloner {
	return &ConanFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Conan, source, target)}
}

func (o *ConanFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ConanFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Conan(sourceRepoDetails)
		})
	}
	return
}

type DockerFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDockerFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *DockerFederatedRepositoryCloner {
	return &DockerFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Docker, source, target)}
}

func (o *DockerFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DockerFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Docker(sourceRepoDetails)
		})
	}
	return
}

type GoFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGoFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GoFederatedRepositoryCloner {
	return &GoFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Go, source, target)}
}

func (o *GoFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GoFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Go(sourceRepoDetails)
		})
	}
	return
}

type NuGetFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNuGetFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *NuGetFederatedRepositoryCloner {
	return &NuGetFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(NuGet, source, target)}
}

func (o *NuGetFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NugetFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Nuget(sourceRepoDetails)
		})
	}
	return
}

type NpmFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewNpmFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *NpmFederatedRepositoryCloner {
	return &NpmFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Npm, source, target)}
}

func (o *NpmFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.NpmFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Npm(sourceRepoDetails)
		})
	}
	return
}

type PhpComposerFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPhpComposerFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PhpComposerFederatedRepositoryCloner {
	return &PhpComposerFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(PhpComposer, source, target)}
}

func (o *PhpComposerFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.ComposerFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Composer(sourceRepoDetails)
		})
	}
	return
}

type PuppetFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPuppetFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PuppetFederatedRepositoryCloner {
	return &PuppetFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Puppet, source, target)}
}

func (o *PuppetFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PuppetFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Puppet(sourceRepoDetails)
		})
	}
	return
}

type PyPiFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewPyPiFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *PyPiFederatedRepositoryCloner {
	return &PyPiFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(PyPi, source, target)}
}

func (o *PyPiFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.PuppetFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Puppet(sourceRepoDetails)
		})
	}
	return
}

type RubyFederatedGemsRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewRubyFederatedGemsRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *RubyFederatedGemsRepositoryCloner {
	return &RubyFederatedGemsRepositoryCloner{
		NewFederatedRepositoryClonerImpl(RubyGems, source, target)}
}

func (o *RubyFederatedGemsRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GemsFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Gems(sourceRepoDetails)
		})
	}
	return
}

type GenericFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGenericFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GenericFederatedRepositoryCloner {
	return &GenericFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Generic, source, target)}
}

func (o *GenericFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GenericFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Generic(sourceRepoDetails)
		})
	}
	return
}

type MavenFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewMavenFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *MavenFederatedRepositoryCloner {
	return &MavenFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Maven, source, target)}
}

func (o *MavenFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.MavenFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Maven(sourceRepoDetails)
		})
	}
	return
}

type HelmFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewHelmFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *HelmFederatedRepositoryCloner {
	return &HelmFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Helm, source, target)}
}

func (o *HelmFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.HelmFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Helm(sourceRepoDetails)
		})
	}
	return
}

type GitLfsFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGitLfsFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GitLfsFederatedRepositoryCloner {
	return &GitLfsFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(GitLfs, source, target)}
}

func (o *GitLfsFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GitlfsFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Gitlfs(sourceRepoDetails)
		})
	}
	return
}

type DebianFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewDebianFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *DebianFederatedRepositoryCloner {
	return &DebianFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(GitLfs, source, target)}
}

func (o *DebianFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.DebianFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Debian(sourceRepoDetails)
		})
	}
	return
}

type YumFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewYumFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *YumFederatedRepositoryCloner {
	return &YumFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(YUM, source, target)}
}

func (o *YumFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.YumFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Yum(sourceRepoDetails)
		})
	}
	return
}

type VagrantFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewVagrantFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *VagrantFederatedRepositoryCloner {
	return &VagrantFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(YUM, source, target)}
}

func (o *VagrantFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.VagrantFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Vagrant(sourceRepoDetails)
		})
	}
	return
}

type CargoFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewCargoFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *CargoFederatedRepositoryCloner {
	return &CargoFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Cargo, source, target)}
}

func (o *CargoFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.CargoFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Cargo(sourceRepoDetails)
		})
	}
	return
}

type GradleFederatedRepositoryCloner struct {
	*RepositoryClonerImpl
}

func NewGradleFederatedRepositoryCloner(
	source *ArtifactoryManager, target *ArtifactoryManager) *GradleFederatedRepositoryCloner {
	return &GradleFederatedRepositoryCloner{
		NewFederatedRepositoryClonerImpl(Gradle, source, target)}
}

func (o *GradleFederatedRepositoryCloner) Clone(repoKey string) (err error) {
	sourceRepoDetails := services.GradleFederatedRepositoryParams{}
	if err = o.Source.GetRepository(repoKey, &sourceRepoDetails); err == nil {
		err = o.Target.Execute(o.buildLabelCreateRepository(repoKey), func() error {
			prepareRepositoryBaseParams(&sourceRepoDetails.RepositoryBaseParams)
			return o.Target.CreateFederatedRepository().Gradle(sourceRepoDetails)
		})
	}
	return
}
