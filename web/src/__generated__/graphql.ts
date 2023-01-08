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
  modulePropagation: ModulePropagation;
  modulePropagationId: Scalars['ID'];
  modulePropagationRequestId: Scalars['String'];
  orgAccount: OrganizationalAccount;
  orgAccountId: Scalars['ID'];
  requestTime: Scalars['Time'];
  stateKey: Scalars['String'];
  status: RequestStatus;
  terraformConfigurationBase64: Scalars['String'];
  terraformVersion: Scalars['String'];
  terraformWorkflowRequestId: Scalars['String'];
};

export type ApplyExecutionRequests = {
  __typename?: 'ApplyExecutionRequests';
  items: Array<Maybe<ApplyExecutionRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
};

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

export enum CloudPlatform {
  Aws = 'aws',
  Azure = 'azure',
  Gcp = 'gcp'
}

export type ModuleAccountAssociation = {
  __typename?: 'ModuleAccountAssociation';
  applyExecutionRequests: ApplyExecutionRequests;
  modulePropagation: ModulePropagation;
  modulePropagationId: Scalars['ID'];
  orgAccount: OrganizationalAccount;
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
  cloudPlatform: CloudPlatform;
  moduleGroupId: Scalars['ID'];
  modulePropagations: ModulePropagations;
  name: Scalars['String'];
  versions: ModuleVersions;
};


export type ModuleGroupModulePropagationsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
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
  driftCheckRequests: ModulePropagationDriftCheckRequests;
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


export type ModulePropagationDriftCheckRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModulePropagationExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModulePropagationModuleAccountAssociationsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type ModulePropagationDriftCheckRequest = {
  __typename?: 'ModulePropagationDriftCheckRequest';
  modulePropagationDriftCheckRequestId: Scalars['ID'];
  modulePropagationId: Scalars['ID'];
  planExecutionRequests: PlanExecutionRequests;
  requestTime: Scalars['Time'];
  status: RequestStatus;
  terraformDriftCheckWorkflowRequests: TerraformDriftCheckWorkflowRequests;
};


export type ModulePropagationDriftCheckRequestPlanExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModulePropagationDriftCheckRequestTerraformDriftCheckWorkflowRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type ModulePropagationDriftCheckRequests = {
  __typename?: 'ModulePropagationDriftCheckRequests';
  items: Array<Maybe<ModulePropagationDriftCheckRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type ModulePropagationExecutionRequest = {
  __typename?: 'ModulePropagationExecutionRequest';
  applyExecutionRequests: ApplyExecutionRequests;
  modulePropagationExecutionRequestId: Scalars['ID'];
  modulePropagationId: Scalars['ID'];
  planExecutionRequests: PlanExecutionRequests;
  requestTime: Scalars['Time'];
  status: RequestStatus;
  terraformExecutionWorkflowRequests: TerraformExecutionWorkflowRequests;
};


export type ModulePropagationExecutionRequestApplyExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModulePropagationExecutionRequestPlanExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModulePropagationExecutionRequestTerraformExecutionWorkflowRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type ModulePropagationExecutionRequests = {
  __typename?: 'ModulePropagationExecutionRequests';
  items: Array<Maybe<ModulePropagationExecutionRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type ModulePropagationUpdate = {
  arguments?: InputMaybe<Array<ArgumentInput>>;
  awsProviderConfigurations?: InputMaybe<Array<AwsProviderConfigurationInput>>;
  description?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
  orgDimensionId?: InputMaybe<Scalars['ID']>;
  orgUnitId?: InputMaybe<Scalars['ID']>;
};

export type ModulePropagations = {
  __typename?: 'ModulePropagations';
  items: Array<Maybe<ModulePropagation>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type ModuleVariable = {
  __typename?: 'ModuleVariable';
  default?: Maybe<Scalars['String']>;
  description: Scalars['String'];
  name: Scalars['String'];
  type: Scalars['String'];
};

export type ModuleVersion = {
  __typename?: 'ModuleVersion';
  moduleGroup: ModuleGroup;
  moduleGroupId: Scalars['ID'];
  modulePropagations: ModulePropagations;
  moduleVersionId: Scalars['ID'];
  name: Scalars['String'];
  remoteSource: Scalars['String'];
  terraformVersion: Scalars['String'];
  variables: Array<Maybe<ModuleVariable>>;
};


export type ModuleVersionModulePropagationsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
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
  createModulePropagationDriftCheckRequest: ModulePropagationDriftCheckRequest;
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
  updateModulePropagation: ModulePropagation;
  updateOrganizationalUnit: OrganizationalUnit;
};


export type MutationCreateModuleGroupArgs = {
  moduleGroup: NewModuleGroup;
};


export type MutationCreateModulePropagationArgs = {
  modulePropagation: NewModulePropagation;
};


export type MutationCreateModulePropagationDriftCheckRequestArgs = {
  modulePropagationDriftCheckRequest: NewModulePropagationDriftCheckRequest;
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


export type MutationUpdateModulePropagationArgs = {
  modulePropagationId: Scalars['ID'];
  update: ModulePropagationUpdate;
};


export type MutationUpdateOrganizationalUnitArgs = {
  orgDimensionId: Scalars['ID'];
  orgUnitId: Scalars['ID'];
  update: OrganizationalUnitUpdate;
};

export type NewModuleGroup = {
  cloudPlatform: CloudPlatform;
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

export type NewModulePropagationDriftCheckRequest = {
  modulePropagationId: Scalars['ID'];
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
  cloudPlatform: CloudPlatform;
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
  cloudPlatform: CloudPlatform;
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
  modulePropagation: ModulePropagation;
  modulePropagationId: Scalars['ID'];
  modulePropagationRequestId: Scalars['String'];
  orgAccount: OrganizationalAccount;
  orgAccountId: Scalars['ID'];
  planExecutionRequestId: Scalars['ID'];
  planOutput?: Maybe<TerraformPlanOutput>;
  requestTime: Scalars['Time'];
  stateKey: Scalars['String'];
  status: RequestStatus;
  terraformConfigurationBase64: Scalars['String'];
  terraformVersion: Scalars['String'];
  terraformWorkflowRequestId: Scalars['String'];
};

export type PlanExecutionRequests = {
  __typename?: 'PlanExecutionRequests';
  items: Array<Maybe<PlanExecutionRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type Query = {
  __typename?: 'Query';
  applyExecutionRequest: ApplyExecutionRequest;
  applyExecutionRequests: ApplyExecutionRequests;
  moduleAccountAssociation: ModuleAccountAssociation;
  moduleAccountAssociations: ModuleAccountAssociations;
  moduleGroup: ModuleGroup;
  moduleGroups: ModuleGroups;
  modulePropagation: ModulePropagation;
  modulePropagationDriftCheckRequest: ModulePropagationDriftCheckRequest;
  modulePropagationDriftCheckRequests: ModulePropagationDriftCheckRequests;
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


export type QueryModulePropagationDriftCheckRequestArgs = {
  modulePropagationDriftCheckRequestId: Scalars['ID'];
  modulePropagationId: Scalars['ID'];
};


export type QueryModulePropagationDriftCheckRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
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

export enum RequestStatus {
  Failed = 'FAILED',
  Pending = 'PENDING',
  Running = 'RUNNING',
  Succeeded = 'SUCCEEDED'
}

export type TerraformApplyOutput = {
  __typename?: 'TerraformApplyOutput';
  Stderr?: Maybe<Scalars['String']>;
  Stdout?: Maybe<Scalars['String']>;
};

export enum TerraformDriftCheckStatus {
  InSync = 'IN_SYNC',
  OutOfSync = 'OUT_OF_SYNC',
  Pending = 'PENDING'
}

export type TerraformDriftCheckWorkflowRequest = {
  __typename?: 'TerraformDriftCheckWorkflowRequest';
  destroy: Scalars['Boolean'];
  moduleAccountAssociation: ModuleAccountAssociation;
  moduleAccountAssociationKey: Scalars['String'];
  modulePropagation: ModulePropagation;
  modulePropagationDriftCheckRequest: ModulePropagationDriftCheckRequest;
  modulePropagationDriftCheckRequestId: Scalars['ID'];
  modulePropagationId: Scalars['ID'];
  orgAccount: OrganizationalAccount;
  orgAccountId: Scalars['ID'];
  planExecutionRequest?: Maybe<PlanExecutionRequest>;
  planExecutionRequestId?: Maybe<Scalars['ID']>;
  requestTime: Scalars['Time'];
  status: RequestStatus;
  syncStatus: TerraformDriftCheckStatus;
  terraformDriftCheckWorkflowRequestId: Scalars['ID'];
};

export type TerraformDriftCheckWorkflowRequests = {
  __typename?: 'TerraformDriftCheckWorkflowRequests';
  items: Array<Maybe<TerraformDriftCheckWorkflowRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type TerraformExecutionWorkflowRequest = {
  __typename?: 'TerraformExecutionWorkflowRequest';
  applyExecutionRequest?: Maybe<ApplyExecutionRequest>;
  applyExecutionRequestId?: Maybe<Scalars['ID']>;
  destroy: Scalars['Boolean'];
  moduleAccountAssociation: ModuleAccountAssociation;
  moduleAccountAssociationKey: Scalars['String'];
  modulePropagation: ModulePropagation;
  modulePropagationExecutionRequest: ModulePropagationExecutionRequest;
  modulePropagationExecutionRequestId: Scalars['ID'];
  modulePropagationId: Scalars['ID'];
  orgAccount: OrganizationalAccount;
  orgAccountId: Scalars['ID'];
  planExecutionRequest?: Maybe<PlanExecutionRequest>;
  planExecutionRequestId?: Maybe<Scalars['ID']>;
  requestTime: Scalars['Time'];
  status: RequestStatus;
  terraformExecutionWorkflowRequestId: Scalars['ID'];
};

export type TerraformExecutionWorkflowRequests = {
  __typename?: 'TerraformExecutionWorkflowRequests';
  items: Array<Maybe<TerraformExecutionWorkflowRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
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


export type ApplyExecutionRequestQuery = { __typename?: 'Query', applyExecutionRequest: { __typename?: 'ApplyExecutionRequest', applyExecutionRequestId: string, status: RequestStatus, requestTime: any, terraformConfigurationBase64: string, initOutput?: { __typename?: 'TerraformInitOutput', Stdout?: string | null, Stderr?: string | null } | null, applyOutput?: { __typename?: 'TerraformApplyOutput', Stdout?: string | null, Stderr?: string | null } | null } };

export type ModuleAccountAssociationQueryVariables = Exact<{
  modulePropagationId: Scalars['ID'];
  orgAccountId: Scalars['ID'];
}>;


export type ModuleAccountAssociationQuery = { __typename?: 'Query', moduleAccountAssociation: { __typename?: 'ModuleAccountAssociation', modulePropagationId: string, status?: ModuleAccountAssociationStatus | null, modulePropagation: { __typename?: 'ModulePropagation', name: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string } }, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string }, planExecutionRequests: { __typename?: 'PlanExecutionRequests', items: Array<{ __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: RequestStatus, requestTime: any, modulePropagationId: string, modulePropagationRequestId: string } | null> }, applyExecutionRequests: { __typename?: 'ApplyExecutionRequests', items: Array<{ __typename?: 'ApplyExecutionRequest', applyExecutionRequestId: string, status: RequestStatus, requestTime: any, modulePropagationId: string, modulePropagationRequestId: string } | null> } } };

export type ModuleGroupsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModuleGroupsQuery = { __typename?: 'Query', moduleGroups: { __typename?: 'ModuleGroups', items: Array<{ __typename?: 'ModuleGroup', moduleGroupId: string, name: string } | null> } };

export type ModuleGroupQueryVariables = Exact<{
  moduleGroupId: Scalars['ID'];
}>;


export type ModuleGroupQuery = { __typename?: 'Query', moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string, versions: { __typename?: 'ModuleVersions', items: Array<{ __typename?: 'ModuleVersion', moduleVersionId: string, name: string, remoteSource: string, terraformVersion: string } | null> }, modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', name: string, modulePropagationId: string, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string }, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string }, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } } | null> } } };

export type ModulePropagationDriftCheckRequestQueryVariables = Exact<{
  modulePropagationId: Scalars['ID'];
  modulePropagationDriftCheckRequestId: Scalars['ID'];
}>;


export type ModulePropagationDriftCheckRequestQuery = { __typename?: 'Query', modulePropagationDriftCheckRequest: { __typename?: 'ModulePropagationDriftCheckRequest', modulePropagationId: string, modulePropagationDriftCheckRequestId: string, requestTime: any, status: RequestStatus, terraformDriftCheckWorkflowRequests: { __typename?: 'TerraformDriftCheckWorkflowRequests', items: Array<{ __typename?: 'TerraformDriftCheckWorkflowRequest', terraformDriftCheckWorkflowRequestId: string, status: RequestStatus, requestTime: any, destroy: boolean, syncStatus: TerraformDriftCheckStatus, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string }, planExecutionRequest?: { __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: RequestStatus, requestTime: any } | null } | null> } } };

export type ModulePropagationExecutionRequestQueryVariables = Exact<{
  modulePropagationId: Scalars['ID'];
  modulePropagationExecutionRequestId: Scalars['ID'];
}>;


export type ModulePropagationExecutionRequestQuery = { __typename?: 'Query', modulePropagationExecutionRequest: { __typename?: 'ModulePropagationExecutionRequest', modulePropagationId: string, modulePropagationExecutionRequestId: string, requestTime: any, status: RequestStatus, terraformExecutionWorkflowRequests: { __typename?: 'TerraformExecutionWorkflowRequests', items: Array<{ __typename?: 'TerraformExecutionWorkflowRequest', terraformExecutionWorkflowRequestId: string, status: RequestStatus, requestTime: any, destroy: boolean, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string }, planExecutionRequest?: { __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: RequestStatus, requestTime: any } | null, applyExecutionRequest?: { __typename?: 'ApplyExecutionRequest', applyExecutionRequestId: string, status: RequestStatus, requestTime: any } | null } | null> } } };

export type ModulePropagationQueryVariables = Exact<{
  modulePropagationId: Scalars['ID'];
}>;


export type ModulePropagationQuery = { __typename?: 'Query', modulePropagation: { __typename?: 'ModulePropagation', modulePropagationId: string, orgUnitId: string, name: string, description: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string }, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, orgDimensionId: string, name: string, downstreamOrgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, orgDimensionId: string, name: string } | null> } }, executionRequests: { __typename?: 'ModulePropagationExecutionRequests', items: Array<{ __typename?: 'ModulePropagationExecutionRequest', modulePropagationId: string, modulePropagationExecutionRequestId: string, requestTime: any, status: RequestStatus } | null> }, driftCheckRequests: { __typename?: 'ModulePropagationDriftCheckRequests', items: Array<{ __typename?: 'ModulePropagationDriftCheckRequest', modulePropagationId: string, modulePropagationDriftCheckRequestId: string, requestTime: any, status: RequestStatus } | null> }, moduleAccountAssociations: { __typename?: 'ModuleAccountAssociations', items: Array<{ __typename?: 'ModuleAccountAssociation', modulePropagationId: string, status?: ModuleAccountAssociationStatus | null, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string } } | null> } } };

export type CreateModulePropagationExecutionRequestMutationVariables = Exact<{
  modulePropagationId: Scalars['ID'];
}>;


export type CreateModulePropagationExecutionRequestMutation = { __typename?: 'Mutation', createModulePropagationExecutionRequest: { __typename?: 'ModulePropagationExecutionRequest', modulePropagationExecutionRequestId: string, status: RequestStatus } };

export type CreateModulePropagationDriftCheckRequestMutationVariables = Exact<{
  modulePropagationId: Scalars['ID'];
}>;


export type CreateModulePropagationDriftCheckRequestMutation = { __typename?: 'Mutation', createModulePropagationDriftCheckRequest: { __typename?: 'ModulePropagationDriftCheckRequest', modulePropagationDriftCheckRequestId: string, status: RequestStatus } };

export type OrgDimensionsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrgDimensionsQuery = { __typename?: 'Query', organizationalDimensions: { __typename?: 'OrganizationalDimensions', items: Array<{ __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string, orgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string } | null> } } | null> } };

export type UpdateModulePropagationMutationVariables = Exact<{
  modulePropagationId: Scalars['ID'];
  update: ModulePropagationUpdate;
}>;


export type UpdateModulePropagationMutation = { __typename?: 'Mutation', updateModulePropagation: { __typename?: 'ModulePropagation', modulePropagationId: string } };

export type ModuleVersionQueryVariables = Exact<{
  moduleGroupId: Scalars['ID'];
  moduleVersionId: Scalars['ID'];
}>;


export type ModuleVersionQuery = { __typename?: 'Query', moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string, remoteSource: string, terraformVersion: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, variables: Array<{ __typename?: 'ModuleVariable', name: string, type: string, description: string, default?: string | null } | null>, modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', name: string, modulePropagationId: string, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string }, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } } | null> } } };

export type OrganizationalAccountQueryVariables = Exact<{
  orgAccountId: Scalars['ID'];
}>;


export type OrganizationalAccountQuery = { __typename?: 'Query', organizationalAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, orgUnitMemberships: { __typename?: 'OrganizationalUnitMemberships', items: Array<{ __typename?: 'OrganizationalUnitMembership', orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string }, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } } | null> }, moduleAccountAssociations: { __typename?: 'ModuleAccountAssociations', items: Array<{ __typename?: 'ModuleAccountAssociation', status?: ModuleAccountAssociationStatus | null, modulePropagationId: string, orgAccountId: string, modulePropagation: { __typename?: 'ModulePropagation', name: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string } } } | null> } } };

export type OrganizationalAccountsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrganizationalAccountsQuery = { __typename?: 'Query', organizationalAccounts: { __typename?: 'OrganizationalAccounts', items: Array<{ __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string } | null> } };

export type OrganizationalDimensionQueryVariables = Exact<{
  orgDimensionId: Scalars['ID'];
}>;


export type OrganizationalDimensionQuery = { __typename?: 'Query', organizationalDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string, orgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, parentOrgUnitId?: string | null, hierarchy: string } | null> }, modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', modulePropagationId: string, moduleGroupId: string, moduleVersionId: string, orgUnitId: string, orgDimensionId: string, name: string, description: string } | null> }, orgUnitMemberships: { __typename?: 'OrganizationalUnitMemberships', items: Array<{ __typename?: 'OrganizationalUnitMembership', orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string }, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string } } | null> } } };

export type OrganizationalDimensionsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrganizationalDimensionsQuery = { __typename?: 'Query', organizationalDimensions: { __typename?: 'OrganizationalDimensions', items: Array<{ __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } | null> } };

export type OrganizationalUnitQueryVariables = Exact<{
  orgUnitId: Scalars['ID'];
  orgDimensionId: Scalars['ID'];
}>;


export type OrganizationalUnitQuery = { __typename?: 'Query', organizationalUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, orgDimensionId: string, name: string, hierarchy: string, parentOrgUnitId?: string | null, children: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, hierarchy: string } | null> }, downstreamOrgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, hierarchy: string } | null> }, orgUnitMemberships: { __typename?: 'OrganizationalUnitMemberships', items: Array<{ __typename?: 'OrganizationalUnitMembership', orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string } } | null> }, modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', modulePropagationId: string, moduleGroupId: string, moduleVersionId: string, orgUnitId: string, orgDimensionId: string, name: string, description: string } | null> } } };

export type PlanExecutionRequestQueryVariables = Exact<{
  planExecutionRequestId: Scalars['ID'];
}>;


export type PlanExecutionRequestQuery = { __typename?: 'Query', planExecutionRequest: { __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: RequestStatus, requestTime: any, terraformConfigurationBase64: string, initOutput?: { __typename?: 'TerraformInitOutput', Stdout?: string | null, Stderr?: string | null } | null, planOutput?: { __typename?: 'TerraformPlanOutput', Stdout?: string | null, Stderr?: string | null } | null } };


export const ApplyExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"applyExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"applyExecutionRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"applyExecutionRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"applyExecutionRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"terraformConfigurationBase64"}},{"kind":"Field","name":{"kind":"Name","value":"initOutput"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Stdout"}},{"kind":"Field","name":{"kind":"Name","value":"Stderr"}}]}},{"kind":"Field","name":{"kind":"Name","value":"applyOutput"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Stdout"}},{"kind":"Field","name":{"kind":"Name","value":"Stderr"}}]}}]}}]}}]} as unknown as DocumentNode<ApplyExecutionRequestQuery, ApplyExecutionRequestQueryVariables>;
export const ModuleAccountAssociationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleAccountAssociation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAccountAssociation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}},{"kind":"Argument","name":{"kind":"Name","value":"orgAccountId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationRequestId"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationRequestId"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleAccountAssociationQuery, ModuleAccountAssociationQueryVariables>;
export const ModuleGroupsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroups"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<ModuleGroupsQuery, ModuleGroupsQueryVariables>;
export const ModuleGroupDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleGroup"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleGroupId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"remoteSource"}},{"kind":"Field","name":{"kind":"Name","value":"terraformVersion"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleGroupQuery, ModuleGroupQueryVariables>;
export const ModulePropagationDriftCheckRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationDriftCheckRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationDriftCheckRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}},{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"terraformDriftCheckWorkflowRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"terraformDriftCheckWorkflowRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"destroy"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}},{"kind":"Field","name":{"kind":"Name","value":"syncStatus"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationDriftCheckRequestQuery, ModulePropagationDriftCheckRequestQueryVariables>;
export const ModulePropagationExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}},{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"terraformExecutionWorkflowRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"terraformExecutionWorkflowRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"destroy"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}},{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationExecutionRequestQuery, ModulePropagationExecutionRequestQueryVariables>;
export const ModulePropagationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"downstreamOrgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"executionRequests"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"5"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"driftCheckRequests"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"5"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAccountAssociations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationQuery, ModulePropagationQueryVariables>;
export const CreateModulePropagationExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModulePropagationExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModulePropagationExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]} as unknown as DocumentNode<CreateModulePropagationExecutionRequestMutation, CreateModulePropagationExecutionRequestMutationVariables>;
export const CreateModulePropagationDriftCheckRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModulePropagationDriftCheckRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModulePropagationDriftCheckRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationDriftCheckRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]} as unknown as DocumentNode<CreateModulePropagationDriftCheckRequestMutation, CreateModulePropagationDriftCheckRequestMutationVariables>;
export const OrgDimensionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"orgDimensions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalDimensions"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"10000"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"10000"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrgDimensionsQuery, OrgDimensionsQueryVariables>;
export const UpdateModulePropagationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateModulePropagation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"update"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ModulePropagationUpdate"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateModulePropagation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}},{"kind":"Argument","name":{"kind":"Name","value":"update"},"value":{"kind":"Variable","name":{"kind":"Name","value":"update"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}}]}}]}}]} as unknown as DocumentNode<UpdateModulePropagationMutation, UpdateModulePropagationMutationVariables>;
export const ModuleVersionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleVersion"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleVersionId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleGroupId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupId"}}},{"kind":"Argument","name":{"kind":"Name","value":"moduleVersionId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleVersionId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"remoteSource"}},{"kind":"Field","name":{"kind":"Name","value":"terraformVersion"}},{"kind":"Field","name":{"kind":"Name","value":"variables"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"default"}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleVersionQuery, ModuleVersionQueryVariables>;
export const OrganizationalAccountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalAccount"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalAccount"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgAccountId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAccountAssociations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalAccountQuery, OrganizationalAccountQueryVariables>;
export const OrganizationalAccountsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalAccounts"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalAccountsQuery, OrganizationalAccountsQueryVariables>;
export const OrganizationalDimensionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalDimension"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalDimension"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgDimensionId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"parentOrgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalDimensionQuery, OrganizationalDimensionQueryVariables>;
export const OrganizationalDimensionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalDimensions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalDimensions"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalDimensionsQuery, OrganizationalDimensionsQueryVariables>;
export const OrganizationalUnitDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalUnit"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalUnit"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgDimensionId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}}},{"kind":"Argument","name":{"kind":"Name","value":"orgUnitId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}},{"kind":"Field","name":{"kind":"Name","value":"parentOrgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"children"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"downstreamOrgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalUnitQuery, OrganizationalUnitQueryVariables>;
export const PlanExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"planExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"planExecutionRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"planExecutionRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"planExecutionRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"terraformConfigurationBase64"}},{"kind":"Field","name":{"kind":"Name","value":"initOutput"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Stdout"}},{"kind":"Field","name":{"kind":"Name","value":"Stderr"}}]}},{"kind":"Field","name":{"kind":"Name","value":"planOutput"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"Stdout"}},{"kind":"Field","name":{"kind":"Name","value":"Stderr"}}]}}]}}]}}]} as unknown as DocumentNode<PlanExecutionRequestQuery, PlanExecutionRequestQueryVariables>;