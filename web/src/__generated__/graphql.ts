/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
};

export type ApplyExecutionRequest = {
  __typename?: 'ApplyExecutionRequest';
  additionalArguments: Array<Maybe<Scalars['String']>>;
  applyExecutionRequestId: Scalars['ID'];
  applyOutput?: Maybe<TerraformApplyOutput>;
  callbackTaskToken: Scalars['String'];
  initOutput?: Maybe<TerraformInitOutput>;
  moduleAccountAssociationKey: Scalars['String'];
  modulePropagationExecutionRequestId: Scalars['String'];
  modulePropagationId: Scalars['ID'];
  requestTime: Scalars['Time'];
  stateKey: Scalars['String'];
  status: ApplyExecutionStatus;
  terraformConfigurationBase64: Scalars['String'];
  terraformVersion: Scalars['String'];
  workflowExecutionId: Scalars['String'];
};

export type ApplyExecutionRequests = {
  __typename?: 'ApplyExecutionRequests';
  items: Array<Maybe<ApplyExecutionRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export enum ApplyExecutionStatus {
  Failed = 'FAILED',
  Pending = 'PENDING',
  Running = 'RUNNING',
  Succeeded = 'SUCCEEDED'
}

export type Argument = {
  __typename?: 'Argument';
  name: Scalars['String'];
  value: Scalars['String'];
};

export type ArgumentInput = {
  name: Scalars['String'];
  value: Scalars['String'];
};

export type AwsProviderConfiguration = {
  __typename?: 'AwsProviderConfiguration';
  alias: Scalars['String'];
  region: Scalars['String'];
};

export type AwsProviderConfigurationInput = {
  alias: Scalars['String'];
  region: Scalars['String'];
};

export type ModuleAccountAssociation = {
  __typename?: 'ModuleAccountAssociation';
  applyExecutionRequests: ApplyExecutionRequests;
  modulePropagation: ModulePropagation;
  modulePropagationId: Scalars['ID'];
  orgAccountId: Scalars['ID'];
  planExecutionRequests: PlanExecutionRequests;
  remoteStateBucket: Scalars['String'];
  remoteStateKey: Scalars['String'];
  status?: Maybe<ModuleAccountAssociationStatus>;
  terraformConfiguration: Scalars['String'];
};


export type ModuleAccountAssociationApplyExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModuleAccountAssociationPlanExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export enum ModuleAccountAssociationStatus {
  Active = 'ACTIVE',
  Inactive = 'INACTIVE'
}

export type ModuleAccountAssociations = {
  __typename?: 'ModuleAccountAssociations';
  items: Array<Maybe<ModuleAccountAssociation>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type ModuleGroup = {
  __typename?: 'ModuleGroup';
  moduleGroupId: Scalars['ID'];
  name: Scalars['String'];
  versions: ModuleVersions;
};


export type ModuleGroupVersionsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type ModuleGroups = {
  __typename?: 'ModuleGroups';
  items: Array<Maybe<ModuleGroup>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type ModulePropagation = {
  __typename?: 'ModulePropagation';
  arguments: Array<Argument>;
  awsProviderConfigurations: Array<AwsProviderConfiguration>;
  description: Scalars['String'];
  executionRequests: ModulePropagationExecutionRequests;
  moduleAccountAssociations: ModuleAccountAssociations;
  moduleGroup: ModuleGroup;
  moduleGroupId: Scalars['ID'];
  modulePropagationId: Scalars['ID'];
  moduleVersion: ModuleVersion;
  moduleVersionId: Scalars['ID'];
  name: Scalars['String'];
  orgDimension: OrganizationalDimension;
  orgDimensionId: Scalars['ID'];
  orgUnit: OrganizationalUnit;
  orgUnitId: Scalars['ID'];
};


export type ModulePropagationExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModulePropagationModuleAccountAssociationsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type ModulePropagationExecutionRequest = {
  __typename?: 'ModulePropagationExecutionRequest';
  applyExecutionRequests: ApplyExecutionRequests;
  modulePropagationExecutionRequestId: Scalars['ID'];
  modulePropagationId: Scalars['ID'];
  planExecutionRequests: PlanExecutionRequests;
  requestTime: Scalars['Time'];
  status: ModulePropagationExecutionRequestStatus;
};


export type ModulePropagationExecutionRequestApplyExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModulePropagationExecutionRequestPlanExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export enum ModulePropagationExecutionRequestStatus {
  Failed = 'FAILED',
  Pending = 'PENDING',
  Running = 'RUNNING',
  Succeeded = 'SUCCEEDED'
}

export type ModulePropagationExecutionRequests = {
  __typename?: 'ModulePropagationExecutionRequests';
  items: Array<Maybe<ModulePropagationExecutionRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type ModulePropagations = {
  __typename?: 'ModulePropagations';
  items: Array<Maybe<ModulePropagation>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type ModuleVariable = {
  __typename?: 'ModuleVariable';
  Default?: Maybe<Scalars['String']>;
  Description: Scalars['String'];
  Name: Scalars['String'];
  Type: Scalars['String'];
};

export type ModuleVersion = {
  __typename?: 'ModuleVersion';
  moduleGroupId: Scalars['ID'];
  moduleVersionId: Scalars['ID'];
  name: Scalars['String'];
  remoteSource: Scalars['String'];
  terraformVersion: Scalars['String'];
  variables: Array<Maybe<ModuleVariable>>;
};

export type ModuleVersions = {
  __typename?: 'ModuleVersions';
  items: Array<Maybe<ModuleVersion>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type Mutation = {
  __typename?: 'Mutation';
  createModuleGroup: ModuleGroup;
  createModulePropagation: ModulePropagation;
  createModulePropagationExecutionRequest: ModulePropagationExecutionRequest;
  createModuleVersion: ModuleVersion;
  createOrganizationalAccount: OrganizationalAccount;
  createOrganizationalDimension: OrganizationalDimension;
  createOrganizationalUnit: OrganizationalUnit;
  createOrganizationalUnitMembership: OrganizationalUnitMembership;
  deleteModuleGroup: Scalars['Boolean'];
  deleteModulePropagation: Scalars['Boolean'];
  deleteModuleVersion: Scalars['Boolean'];
  deleteOrganizationalAccount: Scalars['Boolean'];
  deleteOrganizationalDimension: Scalars['Boolean'];
  deleteOrganizationalUnit: Scalars['Boolean'];
  deleteOrganizationalUnitMembership: Scalars['Boolean'];
  updateOrganizationalUnit: OrganizationalUnit;
};


export type MutationCreateModuleGroupArgs = {
  moduleGroup: NewModuleGroup;
};


export type MutationCreateModulePropagationArgs = {
  modulePropagation: NewModulePropagation;
};


export type MutationCreateModulePropagationExecutionRequestArgs = {
  modulePropagationExecutionRequest: NewModulePropagationExecutionRequest;
};


export type MutationCreateModuleVersionArgs = {
  moduleVersion: NewModuleVersion;
};


export type MutationCreateOrganizationalAccountArgs = {
  orgAccount: NewOrganizationalAccount;
};


export type MutationCreateOrganizationalDimensionArgs = {
  orgDimension: NewOrganizationalDimension;
};


export type MutationCreateOrganizationalUnitArgs = {
  orgUnit: NewOrganizationalUnit;
};


export type MutationCreateOrganizationalUnitMembershipArgs = {
  orgUnitMembership: NewOrganizationalUnitMembership;
};


export type MutationDeleteModuleGroupArgs = {
  moduleGroupId: Scalars['ID'];
};


export type MutationDeleteModulePropagationArgs = {
  modulePropagationId: Scalars['ID'];
};


export type MutationDeleteModuleVersionArgs = {
  moduleGroupId: Scalars['ID'];
  moduleVersionId: Scalars['ID'];
};


export type MutationDeleteOrganizationalAccountArgs = {
  orgAccountId: Scalars['ID'];
};


export type MutationDeleteOrganizationalDimensionArgs = {
  orgDimensionId: Scalars['ID'];
};


export type MutationDeleteOrganizationalUnitArgs = {
  orgDimensionId: Scalars['ID'];
  orgUnitId: Scalars['ID'];
};


export type MutationDeleteOrganizationalUnitMembershipArgs = {
  orgAccountId: Scalars['ID'];
  orgDimensionId: Scalars['ID'];
};


export type MutationUpdateOrganizationalUnitArgs = {
  orgDimensionId: Scalars['ID'];
  orgUnitId: Scalars['ID'];
  update: OrganizationalUnitUpdate;
};

export type NewModuleGroup = {
  name: Scalars['String'];
};

export type NewModulePropagation = {
  arguments: Array<ArgumentInput>;
  awsProviderConfigurations: Array<AwsProviderConfigurationInput>;
  description: Scalars['String'];
  moduleGroupId: Scalars['ID'];
  moduleVersionId: Scalars['ID'];
  name: Scalars['String'];
  orgDimensionId: Scalars['ID'];
  orgUnitId: Scalars['ID'];
};

export type NewModulePropagationExecutionRequest = {
  modulePropagationId: Scalars['ID'];
};

export type NewModuleVersion = {
  moduleGroupId: Scalars['ID'];
  name: Scalars['String'];
  remoteSource: Scalars['String'];
  terraformVersion: Scalars['String'];
};

export type NewOrganizationalAccount = {
  cloudIdentifier: Scalars['String'];
  cloudPlatform: Scalars['String'];
  name: Scalars['String'];
};

export type NewOrganizationalDimension = {
  name: Scalars['String'];
};

export type NewOrganizationalUnit = {
  name: Scalars['String'];
  orgDimensionId: Scalars['String'];
  parentOrgUnitId: Scalars['ID'];
};

export type NewOrganizationalUnitMembership = {
  orgAccountId: Scalars['ID'];
  orgDimensionId: Scalars['ID'];
  orgUnitId: Scalars['ID'];
};

export type OrganizationalAccount = {
  __typename?: 'OrganizationalAccount';
  cloudIdentifier: Scalars['String'];
  cloudPlatform: Scalars['String'];
  moduleAccountAssociations: ModuleAccountAssociations;
  name: Scalars['String'];
  orgAccountId: Scalars['ID'];
  orgUnitMemberships: OrganizationalUnitMemberships;
};


export type OrganizationalAccountModuleAccountAssociationsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type OrganizationalAccountOrgUnitMembershipsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type OrganizationalAccounts = {
  __typename?: 'OrganizationalAccounts';
  items: Array<Maybe<OrganizationalAccount>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type OrganizationalDimension = {
  __typename?: 'OrganizationalDimension';
  modulePropagations: ModulePropagations;
  name: Scalars['String'];
  orgDimensionId: Scalars['ID'];
  orgUnitMemberships: OrganizationalUnitMemberships;
  orgUnits: OrganizationalUnits;
  rootOrgUnit: OrganizationalUnit;
  rootOrgUnitId: Scalars['ID'];
};


export type OrganizationalDimensionModulePropagationsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type OrganizationalDimensionOrgUnitMembershipsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type OrganizationalDimensionOrgUnitsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type OrganizationalDimensions = {
  __typename?: 'OrganizationalDimensions';
  items: Array<Maybe<OrganizationalDimension>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type OrganizationalUnit = {
  __typename?: 'OrganizationalUnit';
  children: OrganizationalUnits;
  downstreamOrgUnits: OrganizationalUnits;
  hierarchy: Scalars['String'];
  modulePropagations: ModulePropagations;
  name: Scalars['String'];
  orgDimensionId: Scalars['String'];
  orgUnitId: Scalars['ID'];
  orgUnitMemberships: OrganizationalUnitMemberships;
  parentOrgUnit?: Maybe<OrganizationalUnit>;
  parentOrgUnitId?: Maybe<Scalars['ID']>;
};


export type OrganizationalUnitChildrenArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type OrganizationalUnitDownstreamOrgUnitsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type OrganizationalUnitModulePropagationsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type OrganizationalUnitOrgUnitMembershipsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type OrganizationalUnitMembership = {
  __typename?: 'OrganizationalUnitMembership';
  orgAccount: OrganizationalAccount;
  orgAccountId: Scalars['ID'];
  orgDimension: OrganizationalDimension;
  orgDimensionId: Scalars['ID'];
  orgUnit: OrganizationalUnit;
  orgUnitId: Scalars['ID'];
};

export type OrganizationalUnitMemberships = {
  __typename?: 'OrganizationalUnitMemberships';
  items: Array<Maybe<OrganizationalUnitMembership>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type OrganizationalUnitUpdate = {
  Name?: InputMaybe<Scalars['String']>;
  ParentOrgUnitId?: InputMaybe<Scalars['ID']>;
};

export type OrganizationalUnits = {
  __typename?: 'OrganizationalUnits';
  items: Array<Maybe<OrganizationalUnit>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type PlanExecutionRequest = {
  __typename?: 'PlanExecutionRequest';
  additionalArguments: Array<Maybe<Scalars['String']>>;
  callbackTaskToken: Scalars['String'];
  initOutput?: Maybe<TerraformInitOutput>;
  moduleAccountAssociationKey: Scalars['String'];
  modulePropagationExecutionRequestId: Scalars['String'];
  modulePropagationId: Scalars['ID'];
  planExecutionRequestId: Scalars['ID'];
  planOutput?: Maybe<TerraformPlanOutput>;
  requestTime: Scalars['Time'];
  stateKey: Scalars['String'];
  status: PlanExecutionStatus;
  terraformConfigurationBase64: Scalars['String'];
  terraformVersion: Scalars['String'];
  workflowExecutionId: Scalars['String'];
};

export type PlanExecutionRequests = {
  __typename?: 'PlanExecutionRequests';
  items: Array<Maybe<PlanExecutionRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export enum PlanExecutionStatus {
  Failed = 'FAILED',
  Pending = 'PENDING',
  Running = 'RUNNING',
  Succeeded = 'SUCCEEDED'
}

export type Query = {
  __typename?: 'Query';
  applyExecutionRequest: ApplyExecutionRequest;
  applyExecutionRequests: ApplyExecutionRequests;
  moduleAccountAssociation: ModuleAccountAssociation;
  moduleAccountAssociations: ModuleAccountAssociations;
  moduleGroup: ModuleGroup;
  moduleGroups: ModuleGroups;
  modulePropagation: ModulePropagation;
  modulePropagationExecutionRequest: ModulePropagationExecutionRequest;
  modulePropagationExecutionRequests: ModulePropagationExecutionRequests;
  modulePropagations: ModulePropagations;
  moduleVersion: ModuleVersion;
  moduleVersions: ModuleVersions;
  organizationalAccount: OrganizationalAccount;
  organizationalAccounts: OrganizationalAccounts;
  organizationalDimension: OrganizationalDimension;
  organizationalDimensions: OrganizationalDimensions;
  organizationalUnit: OrganizationalUnit;
  organizationalUnitMembershipsByOrgAccount: OrganizationalUnitMemberships;
  organizationalUnitMembershipsByOrgDimension: OrganizationalUnitMemberships;
  organizationalUnitMembershipsByOrgUnit: OrganizationalUnitMemberships;
  organizationalUnits: OrganizationalUnits;
  organizationalUnitsByDimension: OrganizationalUnits;
  organizationalUnitsByHierarchy: OrganizationalUnits;
  organizationalUnitsByParent: OrganizationalUnits;
  planExecutionRequest: PlanExecutionRequest;
  planExecutionRequests: PlanExecutionRequests;
};


export type QueryApplyExecutionRequestArgs = {
  applyExecutionRequestId: Scalars['ID'];
};


export type QueryApplyExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type QueryModuleAccountAssociationArgs = {
  modulePropagationId: Scalars['ID'];
  orgAccountId: Scalars['ID'];
};


export type QueryModuleAccountAssociationsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type QueryModuleGroupArgs = {
  moduleGroupId: Scalars['ID'];
};


export type QueryModuleGroupsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type QueryModulePropagationArgs = {
  modulePropagationId: Scalars['ID'];
};


export type QueryModulePropagationExecutionRequestArgs = {
  modulePropagationExecutionRequestId: Scalars['ID'];
  modulePropagationId: Scalars['ID'];
};


export type QueryModulePropagationExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type QueryModulePropagationsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type QueryModuleVersionArgs = {
  moduleGroupId: Scalars['ID'];
  moduleVersionId: Scalars['ID'];
};


export type QueryModuleVersionsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  moduleGroupId: Scalars['ID'];
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type QueryOrganizationalAccountArgs = {
  orgAccountId: Scalars['ID'];
};


export type QueryOrganizationalAccountsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type QueryOrganizationalDimensionArgs = {
  orgDimensionId: Scalars['ID'];
};


export type QueryOrganizationalDimensionsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type QueryOrganizationalUnitArgs = {
  orgDimensionId: Scalars['ID'];
  orgUnitId: Scalars['ID'];
};


export type QueryOrganizationalUnitMembershipsByOrgAccountArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
  orgAccountId: Scalars['ID'];
};


export type QueryOrganizationalUnitMembershipsByOrgDimensionArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
  orgDimensionId: Scalars['ID'];
};


export type QueryOrganizationalUnitMembershipsByOrgUnitArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
  orgUnitId: Scalars['ID'];
};


export type QueryOrganizationalUnitsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type QueryOrganizationalUnitsByDimensionArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
  orgDimensionId: Scalars['ID'];
};


export type QueryOrganizationalUnitsByHierarchyArgs = {
  hierarchy: Scalars['String'];
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
  orgDimensionId: Scalars['ID'];
};


export type QueryOrganizationalUnitsByParentArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
  orgDimensionId: Scalars['ID'];
  parentOrgUnitId: Scalars['ID'];
};


export type QueryPlanExecutionRequestArgs = {
  planExecutionRequestId: Scalars['ID'];
};


export type QueryPlanExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type TerraformApplyOutput = {
  __typename?: 'TerraformApplyOutput';
  Stderr?: Maybe<Scalars['String']>;
  Stdout?: Maybe<Scalars['String']>;
};

export type TerraformInitOutput = {
  __typename?: 'TerraformInitOutput';
  Stderr?: Maybe<Scalars['String']>;
  Stdout?: Maybe<Scalars['String']>;
};

export type TerraformPlanOutput = {
  __typename?: 'TerraformPlanOutput';
  PlanFile?: Maybe<Scalars['String']>;
  PlanJSON?: Maybe<Scalars['String']>;
  Stderr?: Maybe<Scalars['String']>;
  Stdout?: Maybe<Scalars['String']>;
};

export type ApplyExecutionRequestQueryVariables = Exact<{
  applyExecutionRequestId: Scalars['ID'];
}>;


export type ApplyExecutionRequestQuery = { __typename?: 'Query', applyExecutionRequest: { __typename?: 'ApplyExecutionRequest', applyExecutionRequestId: string, status: ApplyExecutionStatus, requestTime: any, terraformConfigurationBase64: string, initOutput?: { __typename?: 'TerraformInitOutput', Stdout?: string | null, Stderr?: string | null } | null, applyOutput?: { __typename?: 'TerraformApplyOutput', Stdout?: string | null, Stderr?: string | null } | null } };

export type ModuleAccountAssociationQueryVariables = Exact<{
  modulePropagationId: Scalars['ID'];
  orgAccountId: Scalars['ID'];
}>;


export type ModuleAccountAssociationQuery = { __typename?: 'Query', moduleAccountAssociation: { __typename?: 'ModuleAccountAssociation', modulePropagationId: string, orgAccountId: string, status?: ModuleAccountAssociationStatus | null, modulePropagation: { __typename?: 'ModulePropagation', moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string } }, planExecutionRequests: { __typename?: 'PlanExecutionRequests', items: Array<{ __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: PlanExecutionStatus, requestTime: any, modulePropagationExecutionRequestId: string } | null> }, applyExecutionRequests: { __typename?: 'ApplyExecutionRequests', items: Array<{ __typename?: 'ApplyExecutionRequest', applyExecutionRequestId: string, status: ApplyExecutionStatus, requestTime: any, modulePropagationExecutionRequestId: string } | null> } } };

export type ModulePropagationExecutionRequestQueryVariables = Exact<{
  modulePropagationId: Scalars['ID'];
  modulePropagationExecutionRequestId: Scalars['ID'];
}>;


export type ModulePropagationExecutionRequestQuery = { __typename?: 'Query', modulePropagationExecutionRequest: { __typename?: 'ModulePropagationExecutionRequest', modulePropagationId: string, modulePropagationExecutionRequestId: string, requestTime: any, status: ModulePropagationExecutionRequestStatus, planExecutionRequests: { __typename?: 'PlanExecutionRequests', items: Array<{ __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: PlanExecutionStatus, requestTime: any } | null> }, applyExecutionRequests: { __typename?: 'ApplyExecutionRequests', items: Array<{ __typename?: 'ApplyExecutionRequest', applyExecutionRequestId: string, status: ApplyExecutionStatus, requestTime: any } | null> } } };

export type ModulePropagationQueryVariables = Exact<{
  modulePropagationId: Scalars['ID'];
}>;


export type ModulePropagationQuery = { __typename?: 'Query', modulePropagation: { __typename?: 'ModulePropagation', modulePropagationId: string, orgUnitId: string, name: string, description: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string }, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, orgDimensionId: string, name: string, downstreamOrgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, orgDimensionId: string, name: string } | null> } }, executionRequests: { __typename?: 'ModulePropagationExecutionRequests', items: Array<{ __typename?: 'ModulePropagationExecutionRequest', modulePropagationId: string, modulePropagationExecutionRequestId: string, requestTime: any, status: ModulePropagationExecutionRequestStatus } | null> }, moduleAccountAssociations: { __typename?: 'ModuleAccountAssociations', items: Array<{ __typename?: 'ModuleAccountAssociation', modulePropagationId: string, orgAccountId: string, status?: ModuleAccountAssociationStatus | null } | null> } } };

export type CreateModulePropagationExecutionRequestMutationVariables = Exact<{
  modulePropagationId: Scalars['ID'];
}>;


export type CreateModulePropagationExecutionRequestMutation = { __typename?: 'Mutation', createModulePropagationExecutionRequest: { __typename?: 'ModulePropagationExecutionRequest', modulePropagationExecutionRequestId: string, status: ModulePropagationExecutionRequestStatus } };

export type OrganizationalAccountQueryVariables = Exact<{
  orgAccountId: Scalars['ID'];
}>;


export type OrganizationalAccountQuery = { __typename?: 'Query', organizationalAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, orgUnitMemberships: { __typename?: 'OrganizationalUnitMemberships', items: Array<{ __typename?: 'OrganizationalUnitMembership', orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string }, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } } | null> }, moduleAccountAssociations: { __typename?: 'ModuleAccountAssociations', items: Array<{ __typename?: 'ModuleAccountAssociation', status?: ModuleAccountAssociationStatus | null, modulePropagationId: string, orgAccountId: string, modulePropagation: { __typename?: 'ModulePropagation', moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string } } } | null> } } };

export type OrganizationalAccountsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrganizationalAccountsQuery = { __typename?: 'Query', organizationalAccounts: { __typename?: 'OrganizationalAccounts', items: Array<{ __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: string, cloudIdentifier: string } | null> } };

export type OrganizationalDimensionQueryVariables = Exact<{
  orgDimensionId: Scalars['ID'];
}>;


export type OrganizationalDimensionQuery = { __typename?: 'Query', organizationalDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string, orgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, parentOrgUnitId?: string | null, hierarchy: string } | null> }, modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', modulePropagationId: string, moduleGroupId: string, moduleVersionId: string, orgUnitId: string, orgDimensionId: string, name: string, description: string } | null> }, orgUnitMemberships: { __typename?: 'OrganizationalUnitMemberships', items: Array<{ __typename?: 'OrganizationalUnitMembership', orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: string, cloudIdentifier: string }, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string } } | null> } } };

export type OrganizationalDimensionsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrganizationalDimensionsQuery = { __typename?: 'Query', organizationalDimensions: { __typename?: 'OrganizationalDimensions', items: Array<{ __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } | null> } };

export type OrganizationalUnitQueryVariables = Exact<{
  orgUnitId: Scalars['ID'];
  orgDimensionId: Scalars['ID'];
}>;


export type OrganizationalUnitQuery = { __typename?: 'Query', organizationalUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, orgDimensionId: string, name: string, hierarchy: string, parentOrgUnitId?: string | null, children: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, hierarchy: string } | null> }, downstreamOrgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, hierarchy: string } | null> }, orgUnitMemberships: { __typename?: 'OrganizationalUnitMemberships', items: Array<{ __typename?: 'OrganizationalUnitMembership', orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: string, cloudIdentifier: string } } | null> }, modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', modulePropagationId: string, moduleGroupId: string, moduleVersionId: string, orgUnitId: string, orgDimensionId: string, name: string, description: string } | null> } } };

export type PlanExecutionRequestQueryVariables = Exact<{
  planExecutionRequestId: Scalars['ID'];
}>;


export type PlanExecutionRequestQuery = { __typename?: 'Query', planExecutionRequest: { __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: PlanExecutionStatus, requestTime: any, terraformConfigurationBase64: string, initOutput?: { __typename?: 'TerraformInitOutput', Stdout?: string | null, Stderr?: string | null } | null, planOutput?: { __typename?: 'TerraformPlanOutput', Stdout?: string | null, Stderr?: string | null } | null } };


export const ApplyExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"applyExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"applyExecutionRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"applyExecutionRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"applyExecutionRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"terraformConfigurationBase64"}},{"kind":"Field","name":{"kind":"Name","value":"initOutput"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Stdout"}},{"kind":"Field","name":{"kind":"Name","value":"Stderr"}}]}},{"kind":"Field","name":{"kind":"Name","value":"applyOutput"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Stdout"}},{"kind":"Field","name":{"kind":"Name","value":"Stderr"}}]}}]}}]}}]} as unknown as DocumentNode<ApplyExecutionRequestQuery, ApplyExecutionRequestQueryVariables>;
export const ModuleAccountAssociationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleAccountAssociation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAccountAssociation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}},{"kind":"Argument","name":{"kind":"Name","value":"orgAccountId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleAccountAssociationQuery, ModuleAccountAssociationQueryVariables>;
export const ModulePropagationExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}},{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationExecutionRequestQuery, ModulePropagationExecutionRequestQueryVariables>;
export const ModulePropagationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"downstreamOrgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"executionRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAccountAssociations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationQuery, ModulePropagationQueryVariables>;
export const CreateModulePropagationExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModulePropagationExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModulePropagationExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]} as unknown as DocumentNode<CreateModulePropagationExecutionRequestMutation, CreateModulePropagationExecutionRequestMutationVariables>;
export const OrganizationalAccountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalAccount"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalAccount"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgAccountId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAccountAssociations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalAccountQuery, OrganizationalAccountQueryVariables>;
export const OrganizationalAccountsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalAccounts"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalAccountsQuery, OrganizationalAccountsQueryVariables>;
export const OrganizationalDimensionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalDimension"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalDimension"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgDimensionId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"parentOrgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalDimensionQuery, OrganizationalDimensionQueryVariables>;
export const OrganizationalDimensionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalDimensions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalDimensions"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalDimensionsQuery, OrganizationalDimensionsQueryVariables>;
export const OrganizationalUnitDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalUnit"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalUnit"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgDimensionId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}}},{"kind":"Argument","name":{"kind":"Name","value":"orgUnitId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}},{"kind":"Field","name":{"kind":"Name","value":"parentOrgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"downstreamOrgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalUnitQuery, OrganizationalUnitQueryVariables>;
export const PlanExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"planExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"planExecutionRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"planExecutionRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"planExecutionRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"terraformConfigurationBase64"}},{"kind":"Field","name":{"kind":"Name","value":"initOutput"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Stdout"}},{"kind":"Field","name":{"kind":"Name","value":"Stderr"}}]}},{"kind":"Field","name":{"kind":"Name","value":"planOutput"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Stdout"}},{"kind":"Field","name":{"kind":"Name","value":"Stderr"}}]}}]}}]}}]} as unknown as DocumentNode<PlanExecutionRequestQuery, PlanExecutionRequestQueryVariables>;