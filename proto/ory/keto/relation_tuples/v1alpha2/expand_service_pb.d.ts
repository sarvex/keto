// @generated by protoc-gen-es v1.7.2 with parameter "target=dts"
// @generated from file ory/keto/relation_tuples/v1alpha2/expand_service.proto (package ory.keto.relation_tuples.v1alpha2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { RelationTuple, Subject } from "./relation_tuples_pb.js";

/**
 * @generated from enum ory.keto.relation_tuples.v1alpha2.NodeType
 */
export declare enum NodeType {
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

/**
 * The request for an ExpandService.Expand RPC.
 * Expands the given subject set.
 *
 * @generated from message ory.keto.relation_tuples.v1alpha2.ExpandRequest
 */
export declare class ExpandRequest extends Message<ExpandRequest> {
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
  maxDepth: number;

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
  snaptoken: string;

  constructor(data?: PartialMessage<ExpandRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "ory.keto.relation_tuples.v1alpha2.ExpandRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExpandRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExpandRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExpandRequest;

  static equals(a: ExpandRequest | PlainMessage<ExpandRequest> | undefined, b: ExpandRequest | PlainMessage<ExpandRequest> | undefined): boolean;
}

/**
 * The response for a ExpandService.Expand RPC.
 *
 * @generated from message ory.keto.relation_tuples.v1alpha2.ExpandResponse
 */
export declare class ExpandResponse extends Message<ExpandResponse> {
  /**
   * The tree the requested subject set expands to.
   * The requested subject set is the subject of the root.
   *
   * This field can be nil in some circumstances.
   *
   * @generated from field: ory.keto.relation_tuples.v1alpha2.SubjectTree tree = 1;
   */
  tree?: SubjectTree;

  constructor(data?: PartialMessage<ExpandResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "ory.keto.relation_tuples.v1alpha2.ExpandResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ExpandResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ExpandResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ExpandResponse;

  static equals(a: ExpandResponse | PlainMessage<ExpandResponse> | undefined, b: ExpandResponse | PlainMessage<ExpandResponse> | undefined): boolean;
}

/**
 * @generated from message ory.keto.relation_tuples.v1alpha2.SubjectTree
 */
export declare class SubjectTree extends Message<SubjectTree> {
  /**
   * The type of the node.
   *
   * @generated from field: ory.keto.relation_tuples.v1alpha2.NodeType node_type = 1;
   */
  nodeType: NodeType;

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
  children: SubjectTree[];

  constructor(data?: PartialMessage<SubjectTree>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "ory.keto.relation_tuples.v1alpha2.SubjectTree";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SubjectTree;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SubjectTree;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SubjectTree;

  static equals(a: SubjectTree | PlainMessage<SubjectTree> | undefined, b: SubjectTree | PlainMessage<SubjectTree> | undefined): boolean;
}

