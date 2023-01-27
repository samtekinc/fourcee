/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel-plugin for production.
 */
const documents = {
    "\n  query applyExecutionRequest($applyExecutionRequestID: ID!) {\n    applyExecutionRequest(applyExecutionRequestID: $applyExecutionRequestID) {\n      id\n      status\n      startedAt\n      completedAt\n      terraformConfiguration\n      moduleAssignment {\n        name\n        modulePropagation {\n          name\n        }\n        orgAccount {\n          id\n          name\n          cloudPlatform\n        }\n        moduleGroup {\n          id\n          name\n        }\n        moduleVersion {\n          id\n          name\n        }\n      }\n      initOutput\n      applyOutput\n    }\n  }\n": types.ApplyExecutionRequestDocument,
    "\n  mutation removeAccountFromOrgUnit($orgUnitID: ID!, $orgAccountID: ID!) {\n    removeAccountFromOrgUnit(orgUnitID: $orgUnitID, orgAccountID: $orgAccountID)\n  }\n": types.RemoveAccountFromOrgUnitDocument,
    "\n  query moduleGroups {\n    moduleGroups(limit: 100) {\n      id\n      cloudPlatform\n      name\n      versions {\n        id\n        remoteSource\n        terraformVersion\n        name\n      }\n    }\n  }\n": types.ModuleGroupsDocument,
    "\n  query moduleGroup($moduleGroupID: ID!) {\n    moduleGroup(moduleGroupID: $moduleGroupID) {\n      id\n      name\n      versions {\n        id\n        name\n        remoteSource\n        terraformVersion\n      }\n      modulePropagations {\n        name\n        id\n        moduleVersion {\n          id\n          name\n        }\n        orgUnit {\n          id\n          name\n        }\n        orgDimension {\n          id\n          name\n        }\n      }\n      moduleAssignments {\n        id\n        moduleVersion {\n          id\n          name\n        }\n        modulePropagation {\n          id\n          name\n        }\n        orgAccount {\n          id\n          name\n        }\n        status\n      }\n    }\n  }\n": types.ModuleGroupDocument,
    "\n  query moduleVersion($moduleVersionID: ID!) {\n    moduleVersion(moduleVersionID: $moduleVersionID) {\n      id\n      name\n      moduleGroup {\n        id\n        cloudPlatform\n        name\n      }\n      remoteSource\n      terraformVersion\n      variables {\n        name\n        type\n        description\n        default\n      }\n      modulePropagations {\n        name\n        description\n        id\n        orgUnit {\n          id\n          name\n        }\n        orgDimension {\n          id\n          name\n        }\n      }\n      moduleAssignments {\n        id\n        name\n        description\n        orgAccount {\n          id\n          name\n          cloudPlatform\n        }\n      }\n    }\n  }\n": types.ModuleVersionDocument,
    "\n  mutation createModuleAssignment($moduleAssignment: NewModuleAssignment!) {\n    createModuleAssignment(moduleAssignment: $moduleAssignment) {\n      id\n    }\n  }\n": types.CreateModuleAssignmentDocument,
    "\n  query moduleAssignmentOptions {\n    orgAccounts {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n    }\n    moduleGroups {\n      id\n      name\n      cloudPlatform\n      versions {\n        id\n        name\n        variables {\n          name\n          type\n          default\n          description\n        }\n      }\n    }\n  }\n": types.ModuleAssignmentOptionsDocument,
    "\n  mutation createModuleGroup($moduleGroup: NewModuleGroup!) {\n    createModuleGroup(moduleGroup: $moduleGroup) {\n      id\n    }\n  }\n": types.CreateModuleGroupDocument,
    "\n  mutation createModuleVersion($moduleVersion: NewModuleVersion!) {\n    createModuleVersion(moduleVersion: $moduleVersion) {\n      id\n    }\n  }\n": types.CreateModuleVersionDocument,
    "\n  mutation createOrgAccount($orgAccount: NewOrgAccount!) {\n    createOrgAccount(orgAccount: $orgAccount) {\n      id\n    }\n  }\n": types.CreateOrgAccountDocument,
    "\n  mutation createOrgDimension($orgDimension: NewOrgDimension!) {\n    createOrgDimension(orgDimension: $orgDimension) {\n      id\n    }\n  }\n": types.CreateOrgDimensionDocument,
    "\n  mutation createOrgUnit($orgUnit: NewOrgUnit!) {\n    createOrgUnit(orgUnit: $orgUnit) {\n      id\n    }\n  }\n": types.CreateOrgUnitDocument,
    "\n  mutation addAccountToOrgUnit($orgUnitID: ID!, $orgAccountID: ID!) {\n    addAccountToOrgUnit(orgUnitID: $orgUnitID, orgAccountID: $orgAccountID)\n  }\n": types.AddAccountToOrgUnitDocument,
    "\n  query orgDimensionsAndUnits {\n    orgDimensions {\n      id\n      name\n      orgUnits {\n        id\n        name\n      }\n    }\n  }\n": types.OrgDimensionsAndUnitsDocument,
    "\n  query orgAccountsAndMemberships {\n    orgAccounts {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n      orgUnits {\n        id\n        orgDimensionID\n      }\n    }\n  }\n": types.OrgAccountsAndMembershipsDocument,
    "\n  query orgAccount($orgAccountID: ID!) {\n    orgAccount(orgAccountID: $orgAccountID) {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n      orgUnits {\n        id\n        name\n        orgDimension {\n          id\n          name\n        }\n      }\n      moduleAssignments {\n        id\n        name\n        status\n        moduleGroup {\n          id\n          name\n        }\n        moduleVersion {\n          id\n          name\n        }\n        modulePropagation {\n          id\n          name\n          orgUnit {\n            id\n            name\n          }\n          orgDimension {\n            id\n            name\n          }\n        }\n      }\n    }\n  }\n": types.OrgAccountDocument,
    "\n  query orgAccounts {\n    orgAccounts {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n    }\n  }\n": types.OrgAccountsDocument,
    "\n  query orgDimension($orgDimensionID: ID!) {\n    orgDimension(orgDimensionID: $orgDimensionID) {\n      id\n      name\n      rootOrgUnitID\n      orgUnits {\n        id\n        name\n        parentOrgUnitID\n        hierarchy\n      }\n    }\n  }\n": types.OrgDimensionDocument,
    "\n  query orgDimensions {\n    orgDimensions(limit: 100) {\n      id\n      name\n      orgUnits {\n        id\n      }\n    }\n  }\n": types.OrgDimensionsDocument,
    "\n  query orgUnit($orgUnitID: ID!) {\n    orgUnit(orgUnitID: $orgUnitID) {\n      id\n      name\n\n      orgDimension {\n        id\n        name\n      }\n\n      upstreamOrgUnits {\n        id\n        name\n        modulePropagations {\n          id\n          name\n          description\n          moduleGroup {\n            id\n            name\n          }\n          moduleVersion {\n            id\n            name\n          }\n        }\n      }\n\n      orgAccounts {\n        id\n        name\n        cloudPlatform\n        cloudIdentifier\n      }\n\n      modulePropagations {\n        id\n        name\n        description\n        moduleGroup {\n          id\n          name\n        }\n        moduleVersion {\n          id\n          name\n        }\n      }\n    }\n  }\n": types.OrgUnitDocument,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query applyExecutionRequest($applyExecutionRequestID: ID!) {\n    applyExecutionRequest(applyExecutionRequestID: $applyExecutionRequestID) {\n      id\n      status\n      startedAt\n      completedAt\n      terraformConfiguration\n      moduleAssignment {\n        name\n        modulePropagation {\n          name\n        }\n        orgAccount {\n          id\n          name\n          cloudPlatform\n        }\n        moduleGroup {\n          id\n          name\n        }\n        moduleVersion {\n          id\n          name\n        }\n      }\n      initOutput\n      applyOutput\n    }\n  }\n"): (typeof documents)["\n  query applyExecutionRequest($applyExecutionRequestID: ID!) {\n    applyExecutionRequest(applyExecutionRequestID: $applyExecutionRequestID) {\n      id\n      status\n      startedAt\n      completedAt\n      terraformConfiguration\n      moduleAssignment {\n        name\n        modulePropagation {\n          name\n        }\n        orgAccount {\n          id\n          name\n          cloudPlatform\n        }\n        moduleGroup {\n          id\n          name\n        }\n        moduleVersion {\n          id\n          name\n        }\n      }\n      initOutput\n      applyOutput\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation removeAccountFromOrgUnit($orgUnitID: ID!, $orgAccountID: ID!) {\n    removeAccountFromOrgUnit(orgUnitID: $orgUnitID, orgAccountID: $orgAccountID)\n  }\n"): (typeof documents)["\n  mutation removeAccountFromOrgUnit($orgUnitID: ID!, $orgAccountID: ID!) {\n    removeAccountFromOrgUnit(orgUnitID: $orgUnitID, orgAccountID: $orgAccountID)\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query moduleGroups {\n    moduleGroups(limit: 100) {\n      id\n      cloudPlatform\n      name\n      versions {\n        id\n        remoteSource\n        terraformVersion\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  query moduleGroups {\n    moduleGroups(limit: 100) {\n      id\n      cloudPlatform\n      name\n      versions {\n        id\n        remoteSource\n        terraformVersion\n        name\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query moduleGroup($moduleGroupID: ID!) {\n    moduleGroup(moduleGroupID: $moduleGroupID) {\n      id\n      name\n      versions {\n        id\n        name\n        remoteSource\n        terraformVersion\n      }\n      modulePropagations {\n        name\n        id\n        moduleVersion {\n          id\n          name\n        }\n        orgUnit {\n          id\n          name\n        }\n        orgDimension {\n          id\n          name\n        }\n      }\n      moduleAssignments {\n        id\n        moduleVersion {\n          id\n          name\n        }\n        modulePropagation {\n          id\n          name\n        }\n        orgAccount {\n          id\n          name\n        }\n        status\n      }\n    }\n  }\n"): (typeof documents)["\n  query moduleGroup($moduleGroupID: ID!) {\n    moduleGroup(moduleGroupID: $moduleGroupID) {\n      id\n      name\n      versions {\n        id\n        name\n        remoteSource\n        terraformVersion\n      }\n      modulePropagations {\n        name\n        id\n        moduleVersion {\n          id\n          name\n        }\n        orgUnit {\n          id\n          name\n        }\n        orgDimension {\n          id\n          name\n        }\n      }\n      moduleAssignments {\n        id\n        moduleVersion {\n          id\n          name\n        }\n        modulePropagation {\n          id\n          name\n        }\n        orgAccount {\n          id\n          name\n        }\n        status\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query moduleVersion($moduleVersionID: ID!) {\n    moduleVersion(moduleVersionID: $moduleVersionID) {\n      id\n      name\n      moduleGroup {\n        id\n        cloudPlatform\n        name\n      }\n      remoteSource\n      terraformVersion\n      variables {\n        name\n        type\n        description\n        default\n      }\n      modulePropagations {\n        name\n        description\n        id\n        orgUnit {\n          id\n          name\n        }\n        orgDimension {\n          id\n          name\n        }\n      }\n      moduleAssignments {\n        id\n        name\n        description\n        orgAccount {\n          id\n          name\n          cloudPlatform\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query moduleVersion($moduleVersionID: ID!) {\n    moduleVersion(moduleVersionID: $moduleVersionID) {\n      id\n      name\n      moduleGroup {\n        id\n        cloudPlatform\n        name\n      }\n      remoteSource\n      terraformVersion\n      variables {\n        name\n        type\n        description\n        default\n      }\n      modulePropagations {\n        name\n        description\n        id\n        orgUnit {\n          id\n          name\n        }\n        orgDimension {\n          id\n          name\n        }\n      }\n      moduleAssignments {\n        id\n        name\n        description\n        orgAccount {\n          id\n          name\n          cloudPlatform\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation createModuleAssignment($moduleAssignment: NewModuleAssignment!) {\n    createModuleAssignment(moduleAssignment: $moduleAssignment) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation createModuleAssignment($moduleAssignment: NewModuleAssignment!) {\n    createModuleAssignment(moduleAssignment: $moduleAssignment) {\n      id\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query moduleAssignmentOptions {\n    orgAccounts {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n    }\n    moduleGroups {\n      id\n      name\n      cloudPlatform\n      versions {\n        id\n        name\n        variables {\n          name\n          type\n          default\n          description\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query moduleAssignmentOptions {\n    orgAccounts {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n    }\n    moduleGroups {\n      id\n      name\n      cloudPlatform\n      versions {\n        id\n        name\n        variables {\n          name\n          type\n          default\n          description\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation createModuleGroup($moduleGroup: NewModuleGroup!) {\n    createModuleGroup(moduleGroup: $moduleGroup) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation createModuleGroup($moduleGroup: NewModuleGroup!) {\n    createModuleGroup(moduleGroup: $moduleGroup) {\n      id\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation createModuleVersion($moduleVersion: NewModuleVersion!) {\n    createModuleVersion(moduleVersion: $moduleVersion) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation createModuleVersion($moduleVersion: NewModuleVersion!) {\n    createModuleVersion(moduleVersion: $moduleVersion) {\n      id\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation createOrgAccount($orgAccount: NewOrgAccount!) {\n    createOrgAccount(orgAccount: $orgAccount) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation createOrgAccount($orgAccount: NewOrgAccount!) {\n    createOrgAccount(orgAccount: $orgAccount) {\n      id\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation createOrgDimension($orgDimension: NewOrgDimension!) {\n    createOrgDimension(orgDimension: $orgDimension) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation createOrgDimension($orgDimension: NewOrgDimension!) {\n    createOrgDimension(orgDimension: $orgDimension) {\n      id\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation createOrgUnit($orgUnit: NewOrgUnit!) {\n    createOrgUnit(orgUnit: $orgUnit) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation createOrgUnit($orgUnit: NewOrgUnit!) {\n    createOrgUnit(orgUnit: $orgUnit) {\n      id\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation addAccountToOrgUnit($orgUnitID: ID!, $orgAccountID: ID!) {\n    addAccountToOrgUnit(orgUnitID: $orgUnitID, orgAccountID: $orgAccountID)\n  }\n"): (typeof documents)["\n  mutation addAccountToOrgUnit($orgUnitID: ID!, $orgAccountID: ID!) {\n    addAccountToOrgUnit(orgUnitID: $orgUnitID, orgAccountID: $orgAccountID)\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query orgDimensionsAndUnits {\n    orgDimensions {\n      id\n      name\n      orgUnits {\n        id\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  query orgDimensionsAndUnits {\n    orgDimensions {\n      id\n      name\n      orgUnits {\n        id\n        name\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query orgAccountsAndMemberships {\n    orgAccounts {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n      orgUnits {\n        id\n        orgDimensionID\n      }\n    }\n  }\n"): (typeof documents)["\n  query orgAccountsAndMemberships {\n    orgAccounts {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n      orgUnits {\n        id\n        orgDimensionID\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query orgAccount($orgAccountID: ID!) {\n    orgAccount(orgAccountID: $orgAccountID) {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n      orgUnits {\n        id\n        name\n        orgDimension {\n          id\n          name\n        }\n      }\n      moduleAssignments {\n        id\n        name\n        status\n        moduleGroup {\n          id\n          name\n        }\n        moduleVersion {\n          id\n          name\n        }\n        modulePropagation {\n          id\n          name\n          orgUnit {\n            id\n            name\n          }\n          orgDimension {\n            id\n            name\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query orgAccount($orgAccountID: ID!) {\n    orgAccount(orgAccountID: $orgAccountID) {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n      orgUnits {\n        id\n        name\n        orgDimension {\n          id\n          name\n        }\n      }\n      moduleAssignments {\n        id\n        name\n        status\n        moduleGroup {\n          id\n          name\n        }\n        moduleVersion {\n          id\n          name\n        }\n        modulePropagation {\n          id\n          name\n          orgUnit {\n            id\n            name\n          }\n          orgDimension {\n            id\n            name\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query orgAccounts {\n    orgAccounts {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n    }\n  }\n"): (typeof documents)["\n  query orgAccounts {\n    orgAccounts {\n      id\n      name\n      cloudPlatform\n      cloudIdentifier\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query orgDimension($orgDimensionID: ID!) {\n    orgDimension(orgDimensionID: $orgDimensionID) {\n      id\n      name\n      rootOrgUnitID\n      orgUnits {\n        id\n        name\n        parentOrgUnitID\n        hierarchy\n      }\n    }\n  }\n"): (typeof documents)["\n  query orgDimension($orgDimensionID: ID!) {\n    orgDimension(orgDimensionID: $orgDimensionID) {\n      id\n      name\n      rootOrgUnitID\n      orgUnits {\n        id\n        name\n        parentOrgUnitID\n        hierarchy\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query orgDimensions {\n    orgDimensions(limit: 100) {\n      id\n      name\n      orgUnits {\n        id\n      }\n    }\n  }\n"): (typeof documents)["\n  query orgDimensions {\n    orgDimensions(limit: 100) {\n      id\n      name\n      orgUnits {\n        id\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query orgUnit($orgUnitID: ID!) {\n    orgUnit(orgUnitID: $orgUnitID) {\n      id\n      name\n\n      orgDimension {\n        id\n        name\n      }\n\n      upstreamOrgUnits {\n        id\n        name\n        modulePropagations {\n          id\n          name\n          description\n          moduleGroup {\n            id\n            name\n          }\n          moduleVersion {\n            id\n            name\n          }\n        }\n      }\n\n      orgAccounts {\n        id\n        name\n        cloudPlatform\n        cloudIdentifier\n      }\n\n      modulePropagations {\n        id\n        name\n        description\n        moduleGroup {\n          id\n          name\n        }\n        moduleVersion {\n          id\n          name\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query orgUnit($orgUnitID: ID!) {\n    orgUnit(orgUnitID: $orgUnitID) {\n      id\n      name\n\n      orgDimension {\n        id\n        name\n      }\n\n      upstreamOrgUnits {\n        id\n        name\n        modulePropagations {\n          id\n          name\n          description\n          moduleGroup {\n            id\n            name\n          }\n          moduleVersion {\n            id\n            name\n          }\n        }\n      }\n\n      orgAccounts {\n        id\n        name\n        cloudPlatform\n        cloudIdentifier\n      }\n\n      modulePropagations {\n        id\n        name\n        description\n        moduleGroup {\n          id\n          name\n        }\n        moduleVersion {\n          id\n          name\n        }\n      }\n    }\n  }\n"];

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
**/
export function gql(source: string): unknown;

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;