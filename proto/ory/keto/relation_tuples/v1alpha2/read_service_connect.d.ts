// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=dts"
// @generated from file ory/keto/relation_tuples/v1alpha2/read_service.proto (package ory.keto.relation_tuples.v1alpha2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ListRelationTuplesRequest, ListRelationTuplesResponse } from "./read_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * The service to query relationships.
 *
 * This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis).
 *
 * @generated from service ory.keto.relation_tuples.v1alpha2.ReadService
 */
export declare const ReadService: {
  readonly typeName: "ory.keto.relation_tuples.v1alpha2.ReadService",
  readonly methods: {
    /**
     * Lists ACL relationships.
     *
     * @generated from rpc ory.keto.relation_tuples.v1alpha2.ReadService.ListRelationTuples
     */
    readonly listRelationTuples: {
      readonly name: "ListRelationTuples",
      readonly I: typeof ListRelationTuplesRequest,
      readonly O: typeof ListRelationTuplesResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

