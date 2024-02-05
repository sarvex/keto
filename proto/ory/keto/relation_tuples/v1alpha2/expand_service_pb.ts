// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file ory/keto/relation_tuples/v1alpha2/expand_service.proto (package ory.keto.relation_tuples.v1alpha2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { RelationTuple, Subject } from "./relation_tuples_pb.js";

/**
 * @generated from enum ory.keto.relation_tuples.v1alpha2.NodeType
 */
export enum NodeType {
  /**
   * @generated from enum value: NODE_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * This node expands to a union of all children.
   *
   * @generated from enum value: NODE_TYPE_UNION = 1;
   */
  UNION = 1,

  /**
   * Not implemented yet.
   *
   * @generated from enum value: NODE_TYPE_EXCLUSION = 2;
   */
  EXCLUSION = 2,

  /**
   * Not implemented yet.
   *
   * @generated from enum value: NODE_TYPE_INTERSECTION = 3;
   */
  INTERSECTION = 3,

  /**
   * This node is a leaf and contains no children.
   * Its subject is a `SubjectID` unless `max_depth` was reached.
   *
   * @generated from enum value: NODE_TYPE_LEAF = 4;
   */
  LEAF = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(NodeType)
proto3.util.setEnumType(NodeType, "ory.keto.relation_tuples.v1alpha2.NodeType", [
  { no: 0, name: "NODE_TYPE_UNSPECIFIED" },
  { no: 1, name: "NODE_TYPE_UNION" },
  { no: 2, name: "NODE_TYPE_EXCLUSION" },
  { no: 3, name: "NODE_TYPE_INTERSECTION" },
  { no: 4, name: "NODE_TYPE_LEAF" },
]);

/**
 * The request for an ExpandService.Expand RPC.
 * Expands the given subject set.
 *
 * @generated from message ory.keto.relation_tuples.v1alpha2.ExpandRequest
 */
export class ExpandRequest extends Message<ExpandRequest> {
  /**
   * The subject to expand.
   *
   * @generated from field: ory.keto.relation_tuples.v1alpha2.Subject subject = 1;
   */
  subject?: Subject;

  /**
   * The maximum depth of tree to build.
   *
   * If the value is less than 1 or greater than the global
   * max-depth then the global max-depth will be used instead.
   *
   * It is important to set this parameter to a meaningful
   * value. Ponder how deep you really want to display this.
   *
   * @generated from field: int32 max_depth = 2;
   */
  maxDepth = 0;

  /**
   * This field is not implemented yet and has no effect.
   * <!--
   * Optional. Like reads, a expand is always evaluated at a
   * consistent snapshot no earlier than the given snaptoken.
   *
   * Leave this field blank if you want to expand
   * based on eventually consistent ACLs, benefiting from very
   * low latency, but possibly slightly stale results.
   *
   * If the specified token is too old and no longer known,
   * the server falls back as if no snaptoken had been specified.
   *
   * If not specified the server tries to build the tree
   * on the best snapshot version where it is very likely that
   * ACLs had already been replicated to all availability zones.
   * -->
   *
   * @generated from field: string snaptoken = 3;
   */
  snaptoken = "";

  constructor(data?: PartialMessage<ExpandRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ory.keto.relation_tuples.v1alpha2.ExpandRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "subject", kind: "message", T: Subject },
    { no: 2, name: "max_depth", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "snaptoken", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExpandRequest {
    return new ExpandRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExpandRequest {
    return new ExpandRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExpandRequest {
    return new ExpandRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ExpandRequest | PlainMessage<ExpandRequest> | undefined, b: ExpandRequest | PlainMessage<ExpandRequest> | undefined): boolean {
    return proto3.util.equals(ExpandRequest, a, b);
  }
}

/**
 * The response for a ExpandService.Expand RPC.
 *
 * @generated from message ory.keto.relation_tuples.v1alpha2.ExpandResponse
 */
export class ExpandResponse extends Message<ExpandResponse> {
  /**
   * The tree the requested subject set expands to.
   * The requested subject set is the subject of the root.
   *
   * This field can be nil in some circumstances.
   *
   * @generated from field: ory.keto.relation_tuples.v1alpha2.SubjectTree tree = 1;
   */
  tree?: SubjectTree;

  constructor(data?: PartialMessage<ExpandResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ory.keto.relation_tuples.v1alpha2.ExpandResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "tree", kind: "message", T: SubjectTree },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExpandResponse {
    return new ExpandResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExpandResponse {
    return new ExpandResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExpandResponse {
    return new ExpandResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ExpandResponse | PlainMessage<ExpandResponse> | undefined, b: ExpandResponse | PlainMessage<ExpandResponse> | undefined): boolean {
    return proto3.util.equals(ExpandResponse, a, b);
  }
}

/**
 * @generated from message ory.keto.relation_tuples.v1alpha2.SubjectTree
 */
export class SubjectTree extends Message<SubjectTree> {
  /**
   * The type of the node.
   *
   * @generated from field: ory.keto.relation_tuples.v1alpha2.NodeType node_type = 1;
   */
  nodeType = NodeType.UNSPECIFIED;

  /**
   * The subject this node represents.
   * Deprecated: More information is now available in the tuple field.
   *
   * @generated from field: ory.keto.relation_tuples.v1alpha2.Subject subject = 2 [deprecated = true];
   * @deprecated
   */
  subject?: Subject;

  /**
   * The relation tuple this node represents.
   *
   * @generated from field: ory.keto.relation_tuples.v1alpha2.RelationTuple tuple = 4;
   */
  tuple?: RelationTuple;

  /**
   * The children of this node.
   *
   * This is never set if `node_type` == `NODE_TYPE_LEAF`.
   *
   * @generated from field: repeated ory.keto.relation_tuples.v1alpha2.SubjectTree children = 3;
   */
  children: SubjectTree[] = [];

  constructor(data?: PartialMessage<SubjectTree>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ory.keto.relation_tuples.v1alpha2.SubjectTree";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "node_type", kind: "enum", T: proto3.getEnumType(NodeType) },
    { no: 2, name: "subject", kind: "message", T: Subject },
    { no: 4, name: "tuple", kind: "message", T: RelationTuple },
    { no: 3, name: "children", kind: "message", T: SubjectTree, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubjectTree {
    return new SubjectTree().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubjectTree {
    return new SubjectTree().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubjectTree {
    return new SubjectTree().fromJsonString(jsonString, options);
  }

  static equals(a: SubjectTree | PlainMessage<SubjectTree> | undefined, b: SubjectTree | PlainMessage<SubjectTree> | undefined): boolean {
    return proto3.util.equals(SubjectTree, a, b);
  }
}

