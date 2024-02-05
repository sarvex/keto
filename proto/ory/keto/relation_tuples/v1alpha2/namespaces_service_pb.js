// @generated by protoc-gen-es v1.7.2 with parameter "target=js,js_import_style=legacy_commonjs"
// @generated from file ory/keto/relation_tuples/v1alpha2/namespaces_service.proto (package ory.keto.relation_tuples.v1alpha2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

"use strict";
Object.defineProperty(exports, "__esModule", { value: true });

const { proto3 } = require("@bufbuild/protobuf");

/**
 * Request for ReadService.ListNamespaces RPC.
 *
 * @generated from message ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest
 */
const ListNamespacesRequest = proto3.makeMessageType(
  "ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest",
  [],
);

/**
 * @generated from message ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse
 */
const ListNamespacesResponse = proto3.makeMessageType(
  "ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse",
  () => [
    { no: 1, name: "namespaces", kind: "message", T: Namespace, repeated: true },
  ],
);

/**
 * @generated from message ory.keto.relation_tuples.v1alpha2.Namespace
 */
const Namespace = proto3.makeMessageType(
  "ory.keto.relation_tuples.v1alpha2.Namespace",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);


exports.ListNamespacesRequest = ListNamespacesRequest;
exports.ListNamespacesResponse = ListNamespacesResponse;
exports.Namespace = Namespace;
