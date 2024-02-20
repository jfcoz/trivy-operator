package trivy_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/aquasecurity/trivy-operator/pkg/apis/aquasecurity/v1alpha1"
	"github.com/aquasecurity/trivy-operator/pkg/plugins/trivy"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreateSbomDataSecret(t *testing.T) {
	testCases := []struct {
		name             string
		secretName       string
		sbomDataFilePath string
		wantSecret       corev1.Secret
		err              error
	}{
		{
			name:             "cretae valid sbom data",
			secretName:       "validName",
			sbomDataFilePath: "./testdata/fixture/alpine_sbom.json",
			wantSecret: corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name: "validName",
				},
				Data: map[string][]byte{"bom": []byte(`{"bomFormat":"CycloneDX","specVersion":"1.5","serialNumber":"urn:uuid:9ba1d0c6-b4e3-4bc0-b8f4-2d3d21c7cfc5","version":1,"metadata":{"timestamp":"2023-11-09T23:34:52+00:00","tools":{"components":[{"type":"application","name":"trivy","group":"aquasecurity","version":"0.49.1","supplier":{}}]},"component":{"bom-ref":"pkg:oci/alpine@sha256%3Aeece025e432126ce23f223450a0326fbebde39cdf496a85d8c016293fc851978?arch=amd64\u0026repository_url=index.docker.io%2Flibrary%2Falpine","type":"container","name":"alpine:3.18","purl":"pkg:oci/alpine@sha256%3Aeece025e432126ce23f223450a0326fbebde39cdf496a85d8c016293fc851978?arch=amd64\u0026repository_url=index.docker.io%2Flibrary%2Falpine","supplier":{},"properties":[{"name":"aquasecurity:trivy:DiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:ImageID","value":"sha256:8ca4688f4f356596b5ae539337c9941abc78eda10021d35cbc52659c74d9b443"},{"name":"aquasecurity:trivy:RepoDigest","value":"alpine@sha256:eece025e432126ce23f223450a0326fbebde39cdf496a85d8c016293fc851978"},{"name":"aquasecurity:trivy:RepoTag","value":"alpine:3.18"},{"name":"aquasecurity:trivy:SchemaVersion","value":"2"}]}},"components":[{"bom-ref":"3329179b-b954-4543-87dc-4fd2e651bdec","type":"operating-system","name":"alpine","version":"3.18.4","supplier":{},"properties":[{"name":"aquasecurity:trivy:Class","value":"os-pkgs"},{"name":"aquasecurity:trivy:Type","value":"alpine"}]},{"bom-ref":"pkg:apk/alpine/alpine-baselayout-data@3.4.3-r1?arch=x86_64\u0026distro=3.18.4","type":"library","name":"alpine-baselayout-data","version":"3.4.3-r1","purl":"pkg:apk/alpine/alpine-baselayout-data@3.4.3-r1?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"602007ee374ed96f35e9bf39b1487d67c6afe027"}],"licenses":[{"license":{"name":"GPL-2.0"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"alpine-baselayout-data@3.4.3-r1"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"alpine-baselayout"},{"name":"aquasecurity:trivy:SrcVersion","value":"3.4.3-r1"}]},{"bom-ref":"pkg:apk/alpine/alpine-baselayout@3.4.3-r1?arch=x86_64\u0026distro=3.18.4","type":"library","name":"alpine-baselayout","version":"3.4.3-r1","purl":"pkg:apk/alpine/alpine-baselayout@3.4.3-r1?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"cf0bca32762cd5be9974f4c127467b0f93f78f20"}],"licenses":[{"license":{"name":"GPL-2.0"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"alpine-baselayout@3.4.3-r1"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"alpine-baselayout"},{"name":"aquasecurity:trivy:SrcVersion","value":"3.4.3-r1"}]},{"bom-ref":"pkg:apk/alpine/alpine-keys@2.4-r1?arch=x86_64\u0026distro=3.18.4","type":"library","name":"alpine-keys","version":"2.4-r1","purl":"pkg:apk/alpine/alpine-keys@2.4-r1?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"ec3a3d5ef4c7a168d09516097bb3219ca77c1534"}],"licenses":[{"license":{"name":"MIT"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"alpine-keys@2.4-r1"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"alpine-keys"},{"name":"aquasecurity:trivy:SrcVersion","value":"2.4-r1"}]},{"bom-ref":"pkg:apk/alpine/apk-tools@2.14.0-r2?arch=x86_64\u0026distro=3.18.4","type":"library","name":"apk-tools","version":"2.14.0-r2","purl":"pkg:apk/alpine/apk-tools@2.14.0-r2?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"8cde25f239ebf691cd135a3954e5193c1ac2ae13"}],"licenses":[{"license":{"name":"GPL-2.0"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"apk-tools@2.14.0-r2"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"apk-tools"},{"name":"aquasecurity:trivy:SrcVersion","value":"2.14.0-r2"}]},{"bom-ref":"pkg:apk/alpine/busybox-binsh@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","type":"library","name":"busybox-binsh","version":"1.36.1-r2","purl":"pkg:apk/alpine/busybox-binsh@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"9e0f4ae337ae0115b922df25796870c68af47114"}],"licenses":[{"license":{"name":"GPL-2.0"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"busybox-binsh@1.36.1-r2"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"busybox"},{"name":"aquasecurity:trivy:SrcVersion","value":"1.36.1-r2"}]},{"bom-ref":"pkg:apk/alpine/busybox@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","type":"library","name":"busybox","version":"1.36.1-r2","purl":"pkg:apk/alpine/busybox@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"810fcbdd40674a382415610741a524503b9ba9d2"}],"licenses":[{"license":{"name":"GPL-2.0"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"busybox@1.36.1-r2"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"busybox"},{"name":"aquasecurity:trivy:SrcVersion","value":"1.36.1-r2"}]},{"bom-ref":"pkg:apk/alpine/ca-certificates-bundle@20230506-r0?arch=x86_64\u0026distro=3.18.4","type":"library","name":"ca-certificates-bundle","version":"20230506-r0","purl":"pkg:apk/alpine/ca-certificates-bundle@20230506-r0?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"47f485d08670a9eb21ebf10e70ae65dc43ab6c3d"}],"licenses":[{"license":{"name":"MPL-2.0"}},{"license":{"name":"MIT"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"ca-certificates-bundle@20230506-r0"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"ca-certificates"},{"name":"aquasecurity:trivy:SrcVersion","value":"20230506-r0"}]},{"bom-ref":"pkg:apk/alpine/libc-utils@0.7.2-r5?arch=x86_64\u0026distro=3.18.4","type":"library","name":"libc-utils","version":"0.7.2-r5","purl":"pkg:apk/alpine/libc-utils@0.7.2-r5?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"2e59dafeb8bca0786540846c686f121ae8348a42"}],"licenses":[{"license":{"name":"BSD-2-Clause"}},{"license":{"name":"BSD-3-Clause"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"libc-utils@0.7.2-r5"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"libc-dev"},{"name":"aquasecurity:trivy:SrcVersion","value":"0.7.2-r5"}]},{"bom-ref":"pkg:apk/alpine/libcrypto3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","type":"library","name":"libcrypto3","version":"3.1.3-r0","purl":"pkg:apk/alpine/libcrypto3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"fa12c7857510118cad0c71e2695361574e3ddd3b"}],"licenses":[{"license":{"name":"Apache-2.0"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"libcrypto3@3.1.3-r0"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"openssl"},{"name":"aquasecurity:trivy:SrcVersion","value":"3.1.3-r0"}]},{"bom-ref":"pkg:apk/alpine/libssl3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","type":"library","name":"libssl3","version":"3.1.3-r0","purl":"pkg:apk/alpine/libssl3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"ceb37221d0f02272791d42e583b952031bcb7957"}],"licenses":[{"license":{"name":"Apache-2.0"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"libssl3@3.1.3-r0"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"openssl"},{"name":"aquasecurity:trivy:SrcVersion","value":"3.1.3-r0"}]},{"bom-ref":"pkg:apk/alpine/musl-utils@1.2.4-r1?arch=x86_64\u0026distro=3.18.4","type":"library","name":"musl-utils","version":"1.2.4-r1","purl":"pkg:apk/alpine/musl-utils@1.2.4-r1?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"c78b141d78d68d4cd83f914fbc92f51d60632f53"}],"licenses":[{"license":{"name":"MIT"}},{"license":{"name":"BSD-2-Clause"}},{"license":{"name":"GPL-2.0"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"musl-utils@1.2.4-r1"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"musl"},{"name":"aquasecurity:trivy:SrcVersion","value":"1.2.4-r1"}]},{"bom-ref":"pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4","type":"library","name":"musl","version":"1.2.4-r1","purl":"pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"daa1cb11a76eed0a41bb3f241c1e440c5de6281e"}],"licenses":[{"license":{"name":"MIT"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"musl@1.2.4-r1"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"musl"},{"name":"aquasecurity:trivy:SrcVersion","value":"1.2.4-r1"}]},{"bom-ref":"pkg:apk/alpine/scanelf@1.3.7-r1?arch=x86_64\u0026distro=3.18.4","type":"library","name":"scanelf","version":"1.3.7-r1","purl":"pkg:apk/alpine/scanelf@1.3.7-r1?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"e27abda38faea3635a2db4d50d007751ea280b43"}],"licenses":[{"license":{"name":"GPL-2.0"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"scanelf@1.3.7-r1"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"pax-utils"},{"name":"aquasecurity:trivy:SrcVersion","value":"1.3.7-r1"}]},{"bom-ref":"pkg:apk/alpine/ssl_client@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","type":"library","name":"ssl_client","version":"1.36.1-r2","purl":"pkg:apk/alpine/ssl_client@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"8fa2c75a96af9a716da588f34241fb6a948854e7"}],"licenses":[{"license":{"name":"GPL-2.0"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"ssl_client@1.36.1-r2"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"busybox"},{"name":"aquasecurity:trivy:SrcVersion","value":"1.36.1-r2"}]},{"bom-ref":"pkg:apk/alpine/zlib@1.2.13-r1?arch=x86_64\u0026distro=3.18.4","type":"library","name":"zlib","version":"1.2.13-r1","purl":"pkg:apk/alpine/zlib@1.2.13-r1?arch=x86_64\u0026distro=3.18.4","supplier":{},"hashes":[{"alg":"SHA-1","content":"2656e848992b378aa40dca24af8cde9e97161174"}],"licenses":[{"license":{"name":"Zlib"}}],"properties":[{"name":"aquasecurity:trivy:LayerDiffID","value":"sha256:cc2447e1835a40530975ab80bb1f872fbab0f2a0faecf2ab16fbbb89b3589438"},{"name":"aquasecurity:trivy:LayerDigest","value":"sha256:96526aa774ef0126ad0fe9e9a95764c5fc37f409ab9e97021e7b4775d82bf6fa"},{"name":"aquasecurity:trivy:PkgID","value":"zlib@1.2.13-r1"},{"name":"aquasecurity:trivy:PkgType","value":"alpine"},{"name":"aquasecurity:trivy:SrcName","value":"zlib"},{"name":"aquasecurity:trivy:SrcVersion","value":"1.2.13-r1"}]}],"dependencies":[{"ref":"3329179b-b954-4543-87dc-4fd2e651bdec","dependsOn":["pkg:apk/alpine/alpine-baselayout-data@3.4.3-r1?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/alpine-baselayout@3.4.3-r1?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/alpine-keys@2.4-r1?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/apk-tools@2.14.0-r2?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/busybox-binsh@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/busybox@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/ca-certificates-bundle@20230506-r0?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/libc-utils@0.7.2-r5?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/libcrypto3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/libssl3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/musl-utils@1.2.4-r1?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/scanelf@1.3.7-r1?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/ssl_client@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/zlib@1.2.13-r1?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/alpine-baselayout-data@3.4.3-r1?arch=x86_64\u0026distro=3.18.4","dependsOn":[]},{"ref":"pkg:apk/alpine/alpine-baselayout@3.4.3-r1?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/alpine-baselayout-data@3.4.3-r1?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/busybox-binsh@1.36.1-r2?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/alpine-keys@2.4-r1?arch=x86_64\u0026distro=3.18.4","dependsOn":[]},{"ref":"pkg:apk/alpine/apk-tools@2.14.0-r2?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/ca-certificates-bundle@20230506-r0?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/libcrypto3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/libssl3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/zlib@1.2.13-r1?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/busybox-binsh@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/busybox@1.36.1-r2?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/busybox@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/ca-certificates-bundle@20230506-r0?arch=x86_64\u0026distro=3.18.4","dependsOn":[]},{"ref":"pkg:apk/alpine/libc-utils@0.7.2-r5?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/musl-utils@1.2.4-r1?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/libcrypto3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/libssl3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/libcrypto3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/musl-utils@1.2.4-r1?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/scanelf@1.3.7-r1?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4","dependsOn":[]},{"ref":"pkg:apk/alpine/scanelf@1.3.7-r1?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/ssl_client@1.36.1-r2?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/libcrypto3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/libssl3@3.1.3-r0?arch=x86_64\u0026distro=3.18.4","pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:apk/alpine/zlib@1.2.13-r1?arch=x86_64\u0026distro=3.18.4","dependsOn":["pkg:apk/alpine/musl@1.2.4-r1?arch=x86_64\u0026distro=3.18.4"]},{"ref":"pkg:oci/alpine@sha256%3Aeece025e432126ce23f223450a0326fbebde39cdf496a85d8c016293fc851978?arch=amd64\u0026repository_url=index.docker.io%2Flibrary%2Falpine","dependsOn":["3329179b-b954-4543-87dc-4fd2e651bdec"]}]}`)},
			},
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sbomFile, err := os.ReadFile(tc.sbomDataFilePath)
			assert.NoError(t, err)
			var bom v1alpha1.BOM
			err = json.Unmarshal([]byte(sbomFile), &bom)
			assert.NoError(t, err)
			got, err := trivy.CreateSbomDataAsSecret(bom, tc.secretName)
			if err == nil {
				assert.Equal(t, tc.wantSecret, got)
			}
		})
	}
}

func TestCreateVolumes(t *testing.T) {
	testCases := []struct {
		name      string
		vm        []corev1.VolumeMount
		v         []corev1.Volume
		cName     string
		sn        string
		fn        string
		mountPath string
	}{
		{
			name:      "cretae volumes",
			vm:        []corev1.VolumeMount{},
			v:         []corev1.Volume{},
			sn:        "test",
			cName:     "cname",
			mountPath: "/sbom-cname",
			fn:        "name",
		},
	}
	tc := testCases[0]
	t.Run(tc.name, func(t *testing.T) {
		trivy.CreateVolumeSbomFiles(&tc.vm, &tc.v, &tc.sn, tc.fn, tc.mountPath, tc.cName)
		assert.Equal(t, len(tc.vm), 1)
		assert.Equal(t, len(tc.v), 1)
		assert.Equal(t, tc.vm[0].Name, "sbomvol-cname")
		assert.Equal(t, tc.vm[0].MountPath, "/sbom-cname")
		assert.Equal(t, tc.v[0].Name, "sbomvol-cname")
		assert.Equal(t, tc.v[0].Secret.SecretName, tc.sn)
		assert.Equal(t, tc.v[0].Secret.Items[0].Key, "bom")
		assert.Equal(t, tc.v[0].Secret.Items[0].Path, tc.fn)
	})

}
