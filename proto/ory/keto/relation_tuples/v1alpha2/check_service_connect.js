// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=js,js_import_style=legacy_commonjs"
// @generated from file ory/keto/relation_tuples/v1alpha2/check_service.proto (package ory.keto.relation_tuples.v1alpha2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

"use strict";
Object.defineProperty(exports, "__esModule", { value: true });

const { CheckRequest, CheckResponse } = require("./check_service_pb.js");
const { MethodKind } = require("@bufbuild/protobuf");

/**
 * The service that performs authorization checks
 * based on the stored Access Control Lists.
 *
 * This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis).
 *
 * @generated from service ory.keto.relation_tuples.v1alpha2.CheckService
 */
const CheckService = {
  typeName: "ory.keto.relation_tuples.v1alpha2.CheckService",
  methods: {
    /**
     * Performs an authorization check.
     *
     * @generated from rpc ory.keto.relation_tuples.v1alpha2.CheckService.Check
     */
    check: {
      name: "Check",
      I: CheckRequest,
      O: CheckResponse,
      kind: MethodKind.Unary,
    },
  }
};


exports.CheckService = CheckService;
