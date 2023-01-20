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
  applyOutput?: Maybe<Scalars['String']>;
  initOutput?: Maybe<Scalars['String']>;
  moduleAssignment: ModuleAssignment;
  moduleAssignmentId: Scalars['ID'];
  requestTime: Scalars['Time'];
  status: RequestStatus;
  terraformConfigurationBase64: Scalars['String'];
  terraformPlanBase64: Scalars['String'];
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
  alias?: Maybe<Scalars['String']>;
  region: Scalars['String'];
};

export type AwsProviderConfigurationInput = {
  alias?: InputMaybe<Scalars['String']>;
  region: Scalars['String'];
};

export enum CloudPlatform {
  Aws = 'aws',
  Azure = 'azure',
  Gcp = 'gcp'
}

export type GcpProviderConfiguration = {
  __typename?: 'GcpProviderConfiguration';
  alias?: Maybe<Scalars['String']>;
  region: Scalars['String'];
};

export type GcpProviderConfigurationInput = {
  alias?: InputMaybe<Scalars['String']>;
  region: Scalars['String'];
};

export type Metadata = {
  __typename?: 'Metadata';
  name: Scalars['String'];
  value: Scalars['String'];
};

export type MetadataInput = {
  name: Scalars['String'];
  value: Scalars['String'];
};

export type ModuleAssignment = {
  __typename?: 'ModuleAssignment';
  applyExecutionRequests: ApplyExecutionRequests;
  arguments: Array<Argument>;
  awsProviderConfigurations?: Maybe<Array<AwsProviderConfiguration>>;
  description: Scalars['String'];
  gcpProviderConfigurations?: Maybe<Array<GcpProviderConfiguration>>;
  moduleAssignmentId: Scalars['ID'];
  moduleGroup: ModuleGroup;
  moduleGroupId: Scalars['ID'];
  modulePropagation?: Maybe<ModulePropagation>;
  modulePropagationId?: Maybe<Scalars['ID']>;
  moduleVersion: ModuleVersion;
  moduleVersionId: Scalars['ID'];
  name: Scalars['String'];
  orgAccount: OrganizationalAccount;
  orgAccountId: Scalars['ID'];
  planExecutionRequests: PlanExecutionRequests;
  remoteStateBucket: Scalars['String'];
  remoteStateKey: Scalars['String'];
  remoteStateRegion: Scalars['String'];
  status: ModuleAssignmentStatus;
  terraformConfiguration: Scalars['String'];
  terraformDriftCheckRequests: TerraformDriftCheckRequests;
  terraformExecutionRequests: TerraformExecutionRequests;
};


export type ModuleAssignmentApplyExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModuleAssignmentPlanExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModuleAssignmentTerraformDriftCheckRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type ModuleAssignmentTerraformExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type ModuleAssignmentFiltersInput = {
  isPropagated?: InputMaybe<Scalars['Boolean']>;
};

export enum ModuleAssignmentStatus {
  Active = 'ACTIVE',
  Inactive = 'INACTIVE'
}

export type ModuleAssignmentUpdate = {
  arguments?: InputMaybe<Array<ArgumentInput>>;
  awsProviderConfigurations?: InputMaybe<Array<AwsProviderConfigurationInput>>;
  description?: InputMaybe<Scalars['String']>;
  gcpProviderConfigurations?: InputMaybe<Array<GcpProviderConfigurationInput>>;
  moduleVersionId?: InputMaybe<Scalars['ID']>;
  name?: InputMaybe<Scalars['String']>;
};

export type ModuleAssignments = {
  __typename?: 'ModuleAssignments';
  items: Array<Maybe<ModuleAssignment>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type ModuleGroup = {
  __typename?: 'ModuleGroup';
  cloudPlatform: CloudPlatform;
  moduleAssignments: ModuleAssignments;
  moduleGroupId: Scalars['ID'];
  modulePropagations: ModulePropagations;
  name: Scalars['String'];
  versions: ModuleVersions;
};


export type ModuleGroupModuleAssignmentsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
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
  awsProviderConfigurations?: Maybe<Array<AwsProviderConfiguration>>;
  description: Scalars['String'];
  driftCheckRequests: ModulePropagationDriftCheckRequests;
  executionRequests: ModulePropagationExecutionRequests;
  gcpProviderConfigurations?: Maybe<Array<GcpProviderConfiguration>>;
  moduleAssignments: ModuleAssignments;
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


export type ModulePropagationModuleAssignmentsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type ModulePropagationAssignment = {
  __typename?: 'ModulePropagationAssignment';
  moduleAssignment: ModuleAssignment;
  moduleAssignmentId: Scalars['ID'];
  modulePropagation: ModulePropagation;
  modulePropagationId: Scalars['ID'];
  orgAccount: OrganizationalAccount;
  orgAccountId: Scalars['ID'];
};

export type ModulePropagationAssignments = {
  __typename?: 'ModulePropagationAssignments';
  items: Array<Maybe<ModulePropagationAssignment>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type ModulePropagationDriftCheckRequest = {
  __typename?: 'ModulePropagationDriftCheckRequest';
  modulePropagation: ModulePropagation;
  modulePropagationDriftCheckRequestId: Scalars['ID'];
  modulePropagationId: Scalars['ID'];
  requestTime: Scalars['Time'];
  status: RequestStatus;
  syncStatus: TerraformDriftCheckStatus;
  terraformDriftCheckRequests: TerraformDriftCheckRequests;
};


export type ModulePropagationDriftCheckRequestTerraformDriftCheckRequestsArgs = {
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
  modulePropagation: ModulePropagation;
  modulePropagationExecutionRequestId: Scalars['ID'];
  modulePropagationId: Scalars['ID'];
  requestTime: Scalars['Time'];
  status: RequestStatus;
  terraformExecutionRequests: TerraformExecutionRequests;
};


export type ModulePropagationExecutionRequestTerraformExecutionRequestsArgs = {
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
  gcpProviderConfigurations?: InputMaybe<Array<GcpProviderConfigurationInput>>;
  moduleVersionId?: InputMaybe<Scalars['ID']>;
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
  moduleAssignments: ModuleAssignments;
  moduleGroup: ModuleGroup;
  moduleGroupId: Scalars['ID'];
  modulePropagations: ModulePropagations;
  moduleVersionId: Scalars['ID'];
  name: Scalars['String'];
  remoteSource: Scalars['String'];
  terraformVersion: Scalars['String'];
  variables: Array<Maybe<ModuleVariable>>;
};


export type ModuleVersionModuleAssignmentsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
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
  createModuleAssignment: ModuleAssignment;
  createModuleGroup: ModuleGroup;
  createModulePropagation: ModulePropagation;
  createModulePropagationDriftCheckRequest: ModulePropagationDriftCheckRequest;
  createModulePropagationExecutionRequest: ModulePropagationExecutionRequest;
  createModuleVersion: ModuleVersion;
  createOrganizationalAccount: OrganizationalAccount;
  createOrganizationalDimension: OrganizationalDimension;
  createOrganizationalUnit: OrganizationalUnit;
  createOrganizationalUnitMembership: OrganizationalUnitMembership;
  createTerraformDriftCheckRequest: TerraformDriftCheckRequest;
  createTerraformExecutionRequest: TerraformExecutionRequest;
  deleteModuleGroup: Scalars['Boolean'];
  deleteModulePropagation: Scalars['Boolean'];
  deleteModuleVersion: Scalars['Boolean'];
  deleteOrganizationalAccount: Scalars['Boolean'];
  deleteOrganizationalDimension: Scalars['Boolean'];
  deleteOrganizationalUnit: Scalars['Boolean'];
  deleteOrganizationalUnitMembership: Scalars['Boolean'];
  updateModuleAssignment: ModuleAssignment;
  updateModulePropagation: ModulePropagation;
  updateOrganizationalAccount: OrganizationalAccount;
  updateOrganizationalUnit: OrganizationalUnit;
};


export type MutationCreateModuleAssignmentArgs = {
  moduleAssignment: NewModuleAssignment;
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


export type MutationCreateTerraformDriftCheckRequestArgs = {
  terraformDriftCheckRequest: NewTerraformDriftCheckRequest;
};


export type MutationCreateTerraformExecutionRequestArgs = {
  terraformExecutionRequest: NewTerraformExecutionRequest;
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


export type MutationUpdateModuleAssignmentArgs = {
  moduleAssignmentId: Scalars['ID'];
  moduleAssignmentUpdate: ModuleAssignmentUpdate;
};


export type MutationUpdateModulePropagationArgs = {
  modulePropagationId: Scalars['ID'];
  update: ModulePropagationUpdate;
};


export type MutationUpdateOrganizationalAccountArgs = {
  orgAccountId: Scalars['ID'];
  update: OrganizationalAccountUpdate;
};


export type MutationUpdateOrganizationalUnitArgs = {
  orgDimensionId: Scalars['ID'];
  orgUnitId: Scalars['ID'];
  update: OrganizationalUnitUpdate;
};

export type NewModuleAssignment = {
  arguments: Array<ArgumentInput>;
  awsProviderConfigurations?: InputMaybe<Array<AwsProviderConfigurationInput>>;
  description: Scalars['String'];
  gcpProviderConfigurations?: InputMaybe<Array<GcpProviderConfigurationInput>>;
  moduleGroupId: Scalars['ID'];
  moduleVersionId: Scalars['ID'];
  name: Scalars['String'];
  orgAccountId: Scalars['ID'];
};

export type NewModuleGroup = {
  cloudPlatform: CloudPlatform;
  name: Scalars['String'];
};

export type NewModulePropagation = {
  arguments: Array<ArgumentInput>;
  awsProviderConfigurations?: InputMaybe<Array<AwsProviderConfigurationInput>>;
  description: Scalars['String'];
  gcpProviderConfigurations?: InputMaybe<Array<GcpProviderConfigurationInput>>;
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
  metadata: Array<MetadataInput>;
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

export type NewTerraformDriftCheckRequest = {
  moduleAssignmentId: Scalars['ID'];
};

export type NewTerraformExecutionRequest = {
  destroy: Scalars['Boolean'];
  moduleAssignmentId: Scalars['ID'];
};

export type OrganizationalAccount = {
  __typename?: 'OrganizationalAccount';
  cloudIdentifier: Scalars['String'];
  cloudPlatform: CloudPlatform;
  metadata: Array<Metadata>;
  moduleAssignments: ModuleAssignments;
  name: Scalars['String'];
  orgAccountId: Scalars['ID'];
  orgUnitMemberships: OrganizationalUnitMemberships;
};


export type OrganizationalAccountModuleAssignmentsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};


export type OrganizationalAccountOrgUnitMembershipsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
};

export type OrganizationalAccountUpdate = {
  metadata?: InputMaybe<Array<MetadataInput>>;
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
  orgDimension: OrganizationalDimension;
  orgDimensionId: Scalars['String'];
  orgUnitId: Scalars['ID'];
  orgUnitMemberships: OrganizationalUnitMemberships;
  parentOrgUnit?: Maybe<OrganizationalUnit>;
  parentOrgUnitId?: Maybe<Scalars['ID']>;
  upstreamOrgUnits: OrganizationalUnits;
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
  initOutput?: Maybe<Scalars['String']>;
  moduleAssignment: ModuleAssignment;
  moduleAssignmentId: Scalars['ID'];
  planExecutionRequestId: Scalars['ID'];
  planFile?: Maybe<Scalars['String']>;
  planJSON?: Maybe<Scalars['String']>;
  planOutput?: Maybe<Scalars['String']>;
  requestTime: Scalars['Time'];
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
  moduleAssignment: ModuleAssignment;
  moduleAssignments: ModuleAssignments;
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


export type QueryModuleAssignmentArgs = {
  moduleAssignmentId: Scalars['ID'];
};


export type QueryModuleAssignmentsArgs = {
  filters?: InputMaybe<ModuleAssignmentFiltersInput>;
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
  withOutputs?: InputMaybe<Scalars['Boolean']>;
};


export type QueryPlanExecutionRequestsArgs = {
  limit?: InputMaybe<Scalars['Int']>;
  nextCursor?: InputMaybe<Scalars['String']>;
  withOutputs?: InputMaybe<Scalars['Boolean']>;
};

export enum RequestStatus {
  Failed = 'FAILED',
  Pending = 'PENDING',
  Running = 'RUNNING',
  Succeeded = 'SUCCEEDED'
}

export type TerraformDriftCheckRequest = {
  __typename?: 'TerraformDriftCheckRequest';
  destroy: Scalars['Boolean'];
  moduleAssignment: ModuleAssignment;
  moduleAssignmentId: Scalars['ID'];
  modulePropagation?: Maybe<ModulePropagation>;
  modulePropagationDriftCheckRequest?: Maybe<ModulePropagationDriftCheckRequest>;
  modulePropagationDriftCheckRequestId?: Maybe<Scalars['ID']>;
  modulePropagationId?: Maybe<Scalars['ID']>;
  planExecutionRequest?: Maybe<PlanExecutionRequest>;
  planExecutionRequestId?: Maybe<Scalars['ID']>;
  requestTime: Scalars['Time'];
  status: RequestStatus;
  syncStatus: TerraformDriftCheckStatus;
  terraformDriftCheckRequestId: Scalars['ID'];
};

export type TerraformDriftCheckRequests = {
  __typename?: 'TerraformDriftCheckRequests';
  items: Array<Maybe<TerraformDriftCheckRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export enum TerraformDriftCheckStatus {
  InSync = 'IN_SYNC',
  OutOfSync = 'OUT_OF_SYNC',
  Pending = 'PENDING'
}

export type TerraformExecutionRequest = {
  __typename?: 'TerraformExecutionRequest';
  applyExecutionRequest?: Maybe<ApplyExecutionRequest>;
  applyExecutionRequestId?: Maybe<Scalars['ID']>;
  destroy: Scalars['Boolean'];
  moduleAssignment: ModuleAssignment;
  moduleAssignmentId: Scalars['ID'];
  modulePropagation?: Maybe<ModulePropagation>;
  modulePropagationExecutionRequest?: Maybe<ModulePropagationExecutionRequest>;
  modulePropagationExecutionRequestId?: Maybe<Scalars['ID']>;
  modulePropagationId?: Maybe<Scalars['ID']>;
  planExecutionRequest?: Maybe<PlanExecutionRequest>;
  planExecutionRequestId?: Maybe<Scalars['ID']>;
  requestTime: Scalars['Time'];
  status: RequestStatus;
  terraformExecutionRequestId: Scalars['ID'];
};

export type TerraformExecutionRequests = {
  __typename?: 'TerraformExecutionRequests';
  items: Array<Maybe<TerraformExecutionRequest>>;
  nextCursor?: Maybe<Scalars['String']>;
};

export type ApplyExecutionRequestQueryVariables = Exact<{
  applyExecutionRequestId: Scalars['ID'];
}>;


export type ApplyExecutionRequestQuery = { __typename?: 'Query', applyExecutionRequest: { __typename?: 'ApplyExecutionRequest', applyExecutionRequestId: string, status: RequestStatus, requestTime: any, terraformConfigurationBase64: string, initOutput?: string | null, applyOutput?: string | null, moduleAssignment: { __typename?: 'ModuleAssignment', name: string, moduleAssignmentId: string, modulePropagation?: { __typename?: 'ModulePropagation', modulePropagationId: string, name: string } | null, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform }, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string } } } };

export type DeleteOrganizationalUnitMembershipMutationVariables = Exact<{
  orgDimensionId: Scalars['ID'];
  orgAccountId: Scalars['ID'];
}>;


export type DeleteOrganizationalUnitMembershipMutation = { __typename?: 'Mutation', deleteOrganizationalUnitMembership: boolean };

export type ModuleAssignmentQueryVariables = Exact<{
  moduleAssignmentId: Scalars['ID'];
}>;


export type ModuleAssignmentQuery = { __typename?: 'Query', moduleAssignment: { __typename?: 'ModuleAssignment', name: string, moduleAssignmentId: string, status: ModuleAssignmentStatus, terraformConfiguration: string, modulePropagation?: { __typename?: 'ModulePropagation', modulePropagationId: string, name: string } | null, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform }, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string }, terraformExecutionRequests: { __typename?: 'TerraformExecutionRequests', items: Array<{ __typename?: 'TerraformExecutionRequest', terraformExecutionRequestId: string, status: RequestStatus, requestTime: any, destroy: boolean, moduleAssignment: { __typename?: 'ModuleAssignment', moduleAssignmentId: string, orgAccount: { __typename?: 'OrganizationalAccount', cloudPlatform: CloudPlatform, orgAccountId: string, name: string } }, planExecutionRequest?: { __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: RequestStatus, requestTime: any } | null, applyExecutionRequest?: { __typename?: 'ApplyExecutionRequest', applyExecutionRequestId: string, status: RequestStatus, requestTime: any } | null } | null> }, terraformDriftCheckRequests: { __typename?: 'TerraformDriftCheckRequests', items: Array<{ __typename?: 'TerraformDriftCheckRequest', terraformDriftCheckRequestId: string, status: RequestStatus, requestTime: any, destroy: boolean, syncStatus: TerraformDriftCheckStatus, moduleAssignment: { __typename?: 'ModuleAssignment', moduleAssignmentId: string, orgAccount: { __typename?: 'OrganizationalAccount', cloudPlatform: CloudPlatform, orgAccountId: string, name: string } }, planExecutionRequest?: { __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: RequestStatus, requestTime: any } | null } | null> } } };

export type ModuleAssignmentsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModuleAssignmentsQuery = { __typename?: 'Query', unpropagated: { __typename?: 'ModuleAssignments', items: Array<{ __typename?: 'ModuleAssignment', moduleAssignmentId: string, name: string, description: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string, cloudPlatform: CloudPlatform }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string }, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform } } | null> } };

export type ModuleGroupsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModuleGroupsQuery = { __typename?: 'Query', moduleGroups: { __typename?: 'ModuleGroups', items: Array<{ __typename?: 'ModuleGroup', moduleGroupId: string, cloudPlatform: CloudPlatform, name: string, versions: { __typename?: 'ModuleVersions', items: Array<{ __typename?: 'ModuleVersion', moduleVersionId: string, remoteSource: string, terraformVersion: string, name: string } | null> } } | null> } };

export type ModuleGroupQueryVariables = Exact<{
  moduleGroupId: Scalars['ID'];
}>;


export type ModuleGroupQuery = { __typename?: 'Query', moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, cloudPlatform: CloudPlatform, name: string, versions: { __typename?: 'ModuleVersions', items: Array<{ __typename?: 'ModuleVersion', moduleVersionId: string, name: string, remoteSource: string, terraformVersion: string } | null> }, modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', name: string, modulePropagationId: string, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string }, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string }, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } } | null> }, moduleAssignments: { __typename?: 'ModuleAssignments', items: Array<{ __typename?: 'ModuleAssignment', moduleAssignmentId: string, status: ModuleAssignmentStatus, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string }, modulePropagation?: { __typename?: 'ModulePropagation', modulePropagationId: string, name: string } | null, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string } } | null> } } };

export type ModulePropagationDriftCheckRequestQueryVariables = Exact<{
  modulePropagationId: Scalars['ID'];
  modulePropagationDriftCheckRequestId: Scalars['ID'];
}>;


export type ModulePropagationDriftCheckRequestQuery = { __typename?: 'Query', modulePropagationDriftCheckRequest: { __typename?: 'ModulePropagationDriftCheckRequest', modulePropagationId: string, modulePropagationDriftCheckRequestId: string, requestTime: any, status: RequestStatus, terraformDriftCheckRequests: { __typename?: 'TerraformDriftCheckRequests', items: Array<{ __typename?: 'TerraformDriftCheckRequest', terraformDriftCheckRequestId: string, status: RequestStatus, requestTime: any, destroy: boolean, syncStatus: TerraformDriftCheckStatus, moduleAssignment: { __typename?: 'ModuleAssignment', moduleAssignmentId: string, orgAccount: { __typename?: 'OrganizationalAccount', cloudPlatform: CloudPlatform, orgAccountId: string, name: string } }, planExecutionRequest?: { __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: RequestStatus, requestTime: any } | null } | null> } } };

export type ModulePropagationExecutionRequestQueryVariables = Exact<{
  modulePropagationId: Scalars['ID'];
  modulePropagationExecutionRequestId: Scalars['ID'];
}>;


export type ModulePropagationExecutionRequestQuery = { __typename?: 'Query', modulePropagationExecutionRequest: { __typename?: 'ModulePropagationExecutionRequest', modulePropagationId: string, modulePropagationExecutionRequestId: string, requestTime: any, status: RequestStatus, modulePropagation: { __typename?: 'ModulePropagation', name: string }, terraformExecutionRequests: { __typename?: 'TerraformExecutionRequests', items: Array<{ __typename?: 'TerraformExecutionRequest', terraformExecutionRequestId: string, status: RequestStatus, requestTime: any, destroy: boolean, moduleAssignment: { __typename?: 'ModuleAssignment', moduleAssignmentId: string, orgAccount: { __typename?: 'OrganizationalAccount', cloudPlatform: CloudPlatform, orgAccountId: string, name: string } }, planExecutionRequest?: { __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: RequestStatus, requestTime: any } | null, applyExecutionRequest?: { __typename?: 'ApplyExecutionRequest', applyExecutionRequestId: string, status: RequestStatus, requestTime: any } | null } | null> } } };

export type ModulePropagationQueryVariables = Exact<{
  modulePropagationId: Scalars['ID'];
}>;


export type ModulePropagationQuery = { __typename?: 'Query', modulePropagation: { __typename?: 'ModulePropagation', modulePropagationId: string, orgUnitId: string, name: string, description: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string }, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string }, downstreamOrgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } } | null> } }, executionRequests: { __typename?: 'ModulePropagationExecutionRequests', items: Array<{ __typename?: 'ModulePropagationExecutionRequest', modulePropagationId: string, modulePropagationExecutionRequestId: string, requestTime: any, status: RequestStatus } | null> }, driftCheckRequests: { __typename?: 'ModulePropagationDriftCheckRequests', items: Array<{ __typename?: 'ModulePropagationDriftCheckRequest', modulePropagationId: string, modulePropagationDriftCheckRequestId: string, requestTime: any, status: RequestStatus, syncStatus: TerraformDriftCheckStatus } | null> }, moduleAssignments: { __typename?: 'ModuleAssignments', items: Array<{ __typename?: 'ModuleAssignment', moduleAssignmentId: string, modulePropagationId?: string | null, status: ModuleAssignmentStatus, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string } } | null> } } };

export type ModulePropagationUpdateOptionsQueryVariables = Exact<{
  moduleGroupId: Scalars['ID'];
}>;


export type ModulePropagationUpdateOptionsQuery = { __typename?: 'Query', organizationalDimensions: { __typename?: 'OrganizationalDimensions', items: Array<{ __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string, orgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string } | null> } } | null> }, moduleGroup: { __typename?: 'ModuleGroup', versions: { __typename?: 'ModuleVersions', items: Array<{ __typename?: 'ModuleVersion', moduleVersionId: string, name: string } | null> } } };

export type UpdateModulePropagationMutationVariables = Exact<{
  modulePropagationId: Scalars['ID'];
  update: ModulePropagationUpdate;
}>;


export type UpdateModulePropagationMutation = { __typename?: 'Mutation', updateModulePropagation: { __typename?: 'ModulePropagation', modulePropagationId: string } };

export type ModulePropagationsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModulePropagationsQuery = { __typename?: 'Query', modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', modulePropagationId: string, name: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string, cloudPlatform: CloudPlatform }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string }, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string }, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } } | null> } };

export type ModuleVersionQueryVariables = Exact<{
  moduleGroupId: Scalars['ID'];
  moduleVersionId: Scalars['ID'];
}>;


export type ModuleVersionQuery = { __typename?: 'Query', moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string, remoteSource: string, terraformVersion: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, cloudPlatform: CloudPlatform, name: string }, variables: Array<{ __typename?: 'ModuleVariable', name: string, type: string, description: string, default?: string | null } | null>, modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', name: string, description: string, modulePropagationId: string, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string }, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } } | null> } } };

export type CreateModuleAssignmentMutationVariables = Exact<{
  moduleAssignment: NewModuleAssignment;
}>;


export type CreateModuleAssignmentMutation = { __typename?: 'Mutation', createModuleAssignment: { __typename?: 'ModuleAssignment', moduleAssignmentId: string } };

export type ModuleAssignmentOptionsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModuleAssignmentOptionsQuery = { __typename?: 'Query', organizationalAccounts: { __typename?: 'OrganizationalAccounts', items: Array<{ __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string } | null> }, moduleGroups: { __typename?: 'ModuleGroups', items: Array<{ __typename?: 'ModuleGroup', moduleGroupId: string, name: string, cloudPlatform: CloudPlatform, versions: { __typename?: 'ModuleVersions', items: Array<{ __typename?: 'ModuleVersion', moduleVersionId: string, name: string, variables: Array<{ __typename?: 'ModuleVariable', name: string, type: string, default?: string | null, description: string } | null> } | null> } } | null> } };

export type CreateModuleGroupMutationVariables = Exact<{
  moduleGroup: NewModuleGroup;
}>;


export type CreateModuleGroupMutation = { __typename?: 'Mutation', createModuleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string } };

export type CreateModulePropagationMutationVariables = Exact<{
  modulePropagation: NewModulePropagation;
}>;


export type CreateModulePropagationMutation = { __typename?: 'Mutation', createModulePropagation: { __typename?: 'ModulePropagation', modulePropagationId: string } };

export type ModulePropagationOptionsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModulePropagationOptionsQuery = { __typename?: 'Query', organizationalDimensions: { __typename?: 'OrganizationalDimensions', items: Array<{ __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string, orgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string } | null> } } | null> }, moduleGroups: { __typename?: 'ModuleGroups', items: Array<{ __typename?: 'ModuleGroup', moduleGroupId: string, name: string, cloudPlatform: CloudPlatform, versions: { __typename?: 'ModuleVersions', items: Array<{ __typename?: 'ModuleVersion', moduleVersionId: string, name: string, variables: Array<{ __typename?: 'ModuleVariable', name: string, type: string, default?: string | null, description: string } | null> } | null> } } | null> } };

export type CreateModuleVersionMutationVariables = Exact<{
  moduleVersion: NewModuleVersion;
}>;


export type CreateModuleVersionMutation = { __typename?: 'Mutation', createModuleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string } };

export type CreateOrganizationalAccountMutationVariables = Exact<{
  orgAccount: NewOrganizationalAccount;
}>;


export type CreateOrganizationalAccountMutation = { __typename?: 'Mutation', createOrganizationalAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string } };

export type CreateOrganizationalDimensionMutationVariables = Exact<{
  orgDimension: NewOrganizationalDimension;
}>;


export type CreateOrganizationalDimensionMutation = { __typename?: 'Mutation', createOrganizationalDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string } };

export type CreateOrganizationalUnitMutationVariables = Exact<{
  orgUnit: NewOrganizationalUnit;
}>;


export type CreateOrganizationalUnitMutation = { __typename?: 'Mutation', createOrganizationalUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string } };

export type CreateOrganizationalUnitMembershipMutationVariables = Exact<{
  orgUnitMembership: NewOrganizationalUnitMembership;
}>;


export type CreateOrganizationalUnitMembershipMutation = { __typename?: 'Mutation', createOrganizationalUnitMembership: { __typename?: 'OrganizationalUnitMembership', orgUnitId: string } };

export type OrganizationalDimensionsAndUnitsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrganizationalDimensionsAndUnitsQuery = { __typename?: 'Query', organizationalDimensions: { __typename?: 'OrganizationalDimensions', items: Array<{ __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string, orgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string } | null> } } | null> } };

export type OrganizationalAccountsAndMembershipsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrganizationalAccountsAndMembershipsQuery = { __typename?: 'Query', organizationalAccounts: { __typename?: 'OrganizationalAccounts', items: Array<{ __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string, orgUnitMemberships: { __typename?: 'OrganizationalUnitMemberships', items: Array<{ __typename?: 'OrganizationalUnitMembership', orgDimensionId: string, orgUnitId: string } | null> } } | null> } };

export type OrganizationalAccountQueryVariables = Exact<{
  orgAccountId: Scalars['ID'];
}>;


export type OrganizationalAccountQuery = { __typename?: 'Query', organizationalAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string, orgUnitMemberships: { __typename?: 'OrganizationalUnitMemberships', items: Array<{ __typename?: 'OrganizationalUnitMembership', orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string }, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } } | null> }, moduleAssignments: { __typename?: 'ModuleAssignments', items: Array<{ __typename?: 'ModuleAssignment', moduleAssignmentId: string, status: ModuleAssignmentStatus, modulePropagationId?: string | null, orgAccountId: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string }, modulePropagation?: { __typename?: 'ModulePropagation', modulePropagationId: string, name: string, orgUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string }, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string } } | null } | null> } } };

export type OrganizationalAccountsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrganizationalAccountsQuery = { __typename?: 'Query', organizationalAccounts: { __typename?: 'OrganizationalAccounts', items: Array<{ __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string } | null> } };

export type OrganizationalDimensionQueryVariables = Exact<{
  orgDimensionId: Scalars['ID'];
}>;


export type OrganizationalDimensionQuery = { __typename?: 'Query', organizationalDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string, rootOrgUnitId: string, orgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, parentOrgUnitId?: string | null, hierarchy: string } | null> } } };

export type OrganizationalDimensionsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrganizationalDimensionsQuery = { __typename?: 'Query', organizationalDimensions: { __typename?: 'OrganizationalDimensions', items: Array<{ __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string, orgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string } | null> }, orgUnitMemberships: { __typename?: 'OrganizationalUnitMemberships', items: Array<{ __typename?: 'OrganizationalUnitMembership', orgAccountId: string } | null> } } | null> } };

export type OrganizationalUnitQueryVariables = Exact<{
  orgUnitId: Scalars['ID'];
  orgDimensionId: Scalars['ID'];
}>;


export type OrganizationalUnitQuery = { __typename?: 'Query', organizationalUnit: { __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, hierarchy: string, parentOrgUnitId?: string | null, orgDimension: { __typename?: 'OrganizationalDimension', orgDimensionId: string, name: string }, upstreamOrgUnits: { __typename?: 'OrganizationalUnits', items: Array<{ __typename?: 'OrganizationalUnit', orgUnitId: string, name: string, modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', modulePropagationId: string, name: string, description: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string } } | null> } } | null> }, orgUnitMemberships: { __typename?: 'OrganizationalUnitMemberships', items: Array<{ __typename?: 'OrganizationalUnitMembership', orgDimensionId: string, orgAccountId: string, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string } } | null> }, modulePropagations: { __typename?: 'ModulePropagations', items: Array<{ __typename?: 'ModulePropagation', modulePropagationId: string, name: string, description: string, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string } } | null> } } };

export type PlanExecutionRequestQueryVariables = Exact<{
  planExecutionRequestId: Scalars['ID'];
}>;


export type PlanExecutionRequestQuery = { __typename?: 'Query', planExecutionRequest: { __typename?: 'PlanExecutionRequest', planExecutionRequestId: string, status: RequestStatus, requestTime: any, terraformConfigurationBase64: string, initOutput?: string | null, planOutput?: string | null, moduleAssignment: { __typename?: 'ModuleAssignment', name: string, moduleAssignmentId: string, modulePropagation?: { __typename?: 'ModulePropagation', modulePropagationId: string, name: string } | null, orgAccount: { __typename?: 'OrganizationalAccount', orgAccountId: string, name: string, cloudPlatform: CloudPlatform }, moduleGroup: { __typename?: 'ModuleGroup', moduleGroupId: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', moduleVersionId: string, name: string } } } };

export type CreateModulePropagationDriftCheckRequestMutationVariables = Exact<{
  modulePropagationId: Scalars['ID'];
}>;


export type CreateModulePropagationDriftCheckRequestMutation = { __typename?: 'Mutation', createModulePropagationDriftCheckRequest: { __typename?: 'ModulePropagationDriftCheckRequest', modulePropagationDriftCheckRequestId: string, status: RequestStatus } };

export type CreateModulePropagationExecutionRequestMutationVariables = Exact<{
  modulePropagationId: Scalars['ID'];
}>;


export type CreateModulePropagationExecutionRequestMutation = { __typename?: 'Mutation', createModulePropagationExecutionRequest: { __typename?: 'ModulePropagationExecutionRequest', modulePropagationExecutionRequestId: string, status: RequestStatus } };

export type CreateTerraformDriftCheckRequestMutationVariables = Exact<{
  moduleAssignmentId: Scalars['ID'];
}>;


export type CreateTerraformDriftCheckRequestMutation = { __typename?: 'Mutation', createTerraformDriftCheckRequest: { __typename?: 'TerraformDriftCheckRequest', terraformDriftCheckRequestId: string } };

export type CreateTerraformExecutionRequestMutationVariables = Exact<{
  moduleAssignmentId: Scalars['ID'];
  destroy: Scalars['Boolean'];
}>;


export type CreateTerraformExecutionRequestMutation = { __typename?: 'Mutation', createTerraformExecutionRequest: { __typename?: 'TerraformExecutionRequest', terraformExecutionRequestId: string } };


export const ApplyExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"applyExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"applyExecutionRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"applyExecutionRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"applyExecutionRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"terraformConfigurationBase64"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"initOutput"}},{"kind":"Field","name":{"kind":"Name","value":"applyOutput"}}]}}]}}]} as unknown as DocumentNode<ApplyExecutionRequestQuery, ApplyExecutionRequestQueryVariables>;
export const DeleteOrganizationalUnitMembershipDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"deleteOrganizationalUnitMembership"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"deleteOrganizationalUnitMembership"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgDimensionId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}}},{"kind":"Argument","name":{"kind":"Name","value":"orgAccountId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}}}]}]}}]} as unknown as DocumentNode<DeleteOrganizationalUnitMembershipMutation, DeleteOrganizationalUnitMembershipMutationVariables>;
export const ModuleAssignmentDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleAssignment"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleAssignmentId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"terraformConfiguration"}},{"kind":"Field","name":{"kind":"Name","value":"terraformExecutionRequests"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"5"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"terraformExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"destroy"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}},{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"terraformDriftCheckRequests"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"5"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"terraformDriftCheckRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"destroy"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}},{"kind":"Field","name":{"kind":"Name","value":"syncStatus"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleAssignmentQuery, ModuleAssignmentQueryVariables>;
export const ModuleAssignmentsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleAssignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","alias":{"kind":"Name","value":"unpropagated"},"name":{"kind":"Name","value":"moduleAssignments"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"filters"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"isPropagated"},"value":{"kind":"BooleanValue","value":false}}]}},{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleAssignmentsQuery, ModuleAssignmentsQueryVariables>;
export const ModuleGroupsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroups"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"remoteSource"}},{"kind":"Field","name":{"kind":"Name","value":"terraformVersion"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleGroupsQuery, ModuleGroupsQueryVariables>;
export const ModuleGroupDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleGroup"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleGroupId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"remoteSource"}},{"kind":"Field","name":{"kind":"Name","value":"terraformVersion"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleGroupQuery, ModuleGroupQueryVariables>;
export const ModulePropagationDriftCheckRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationDriftCheckRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationDriftCheckRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}},{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"terraformDriftCheckRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"terraformDriftCheckRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"destroy"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}},{"kind":"Field","name":{"kind":"Name","value":"syncStatus"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationDriftCheckRequestQuery, ModulePropagationDriftCheckRequestQueryVariables>;
export const ModulePropagationExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}},{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"terraformExecutionRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"terraformExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"destroy"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}},{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationExecutionRequestQuery, ModulePropagationExecutionRequestQueryVariables>;
export const ModulePropagationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"downstreamOrgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"executionRequests"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"5"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"driftCheckRequests"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"5"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"syncStatus"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationQuery, ModulePropagationQueryVariables>;
export const ModulePropagationUpdateOptionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationUpdateOptions"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalDimensions"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"10000"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"10000"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleGroupId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationUpdateOptionsQuery, ModulePropagationUpdateOptionsQueryVariables>;
export const UpdateModulePropagationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateModulePropagation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"update"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ModulePropagationUpdate"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateModulePropagation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}},{"kind":"Argument","name":{"kind":"Name","value":"update"},"value":{"kind":"Variable","name":{"kind":"Name","value":"update"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}}]}}]}}]} as unknown as DocumentNode<UpdateModulePropagationMutation, UpdateModulePropagationMutationVariables>;
export const ModulePropagationsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationsQuery, ModulePropagationsQueryVariables>;
export const ModuleVersionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleVersion"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleVersionId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleGroupId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupId"}}},{"kind":"Argument","name":{"kind":"Name","value":"moduleVersionId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleVersionId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"remoteSource"}},{"kind":"Field","name":{"kind":"Name","value":"terraformVersion"}},{"kind":"Field","name":{"kind":"Name","value":"variables"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"default"}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleVersionQuery, ModuleVersionQueryVariables>;
export const CreateModuleAssignmentDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModuleAssignment"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignment"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewModuleAssignment"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModuleAssignment"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleAssignment"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignment"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}}]}}]}}]} as unknown as DocumentNode<CreateModuleAssignmentMutation, CreateModuleAssignmentMutationVariables>;
export const ModuleAssignmentOptionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleAssignmentOptions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"variables"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"default"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleAssignmentOptionsQuery, ModuleAssignmentOptionsQueryVariables>;
export const CreateModuleGroupDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModuleGroup"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroup"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewModuleGroup"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModuleGroup"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleGroup"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroup"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}}]}}]}}]} as unknown as DocumentNode<CreateModuleGroupMutation, CreateModuleGroupMutationVariables>;
export const CreateModulePropagationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModulePropagation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagation"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewModulePropagation"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModulePropagation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagation"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagation"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}}]}}]}}]} as unknown as DocumentNode<CreateModulePropagationMutation, CreateModulePropagationMutationVariables>;
export const ModulePropagationOptionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationOptions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalDimensions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"variables"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"default"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationOptionsQuery, ModulePropagationOptionsQueryVariables>;
export const CreateModuleVersionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModuleVersion"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleVersion"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewModuleVersion"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModuleVersion"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleVersion"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleVersion"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}}]}}]}}]} as unknown as DocumentNode<CreateModuleVersionMutation, CreateModuleVersionMutationVariables>;
export const CreateOrganizationalAccountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createOrganizationalAccount"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccount"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewOrganizationalAccount"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createOrganizationalAccount"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgAccount"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccount"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}}]}}]}}]} as unknown as DocumentNode<CreateOrganizationalAccountMutation, CreateOrganizationalAccountMutationVariables>;
export const CreateOrganizationalDimensionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createOrganizationalDimension"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgDimension"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewOrganizationalDimension"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createOrganizationalDimension"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgDimension"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgDimension"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}}]}}]}}]} as unknown as DocumentNode<CreateOrganizationalDimensionMutation, CreateOrganizationalDimensionMutationVariables>;
export const CreateOrganizationalUnitDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createOrganizationalUnit"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgUnit"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewOrganizationalUnit"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createOrganizationalUnit"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgUnit"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgUnit"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}}]}}]}}]} as unknown as DocumentNode<CreateOrganizationalUnitMutation, CreateOrganizationalUnitMutationVariables>;
export const CreateOrganizationalUnitMembershipDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createOrganizationalUnitMembership"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitMembership"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewOrganizationalUnitMembership"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createOrganizationalUnitMembership"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgUnitMembership"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitMembership"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}}]}}]}}]} as unknown as DocumentNode<CreateOrganizationalUnitMembershipMutation, CreateOrganizationalUnitMembershipMutationVariables>;
export const OrganizationalDimensionsAndUnitsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalDimensionsAndUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalDimensions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalDimensionsAndUnitsQuery, OrganizationalDimensionsAndUnitsQueryVariables>;
export const OrganizationalAccountsAndMembershipsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalAccountsAndMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalAccountsAndMembershipsQuery, OrganizationalAccountsAndMembershipsQueryVariables>;
export const OrganizationalAccountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalAccount"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalAccount"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgAccountId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalAccountQuery, OrganizationalAccountQueryVariables>;
export const OrganizationalAccountsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalAccounts"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalAccountsQuery, OrganizationalAccountsQueryVariables>;
export const OrganizationalDimensionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalDimension"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalDimension"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgDimensionId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"rootOrgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"parentOrgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalDimensionQuery, OrganizationalDimensionQueryVariables>;
export const OrganizationalDimensionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalDimensions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalDimensions"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalDimensionsQuery, OrganizationalDimensionsQueryVariables>;
export const OrganizationalUnitDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"organizationalUnit"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"organizationalUnit"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgDimensionId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionId"}}},{"kind":"Argument","name":{"kind":"Name","value":"orgUnitId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}},{"kind":"Field","name":{"kind":"Name","value":"parentOrgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"upstreamOrgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnitId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnitMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensionId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrganizationalUnitQuery, OrganizationalUnitQueryVariables>;
export const PlanExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"planExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"planExecutionRequestId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"planExecutionRequestId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"planExecutionRequestId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"requestTime"}},{"kind":"Field","name":{"kind":"Name","value":"terraformConfigurationBase64"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignmentId"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccountId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroupId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersionId"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"initOutput"}},{"kind":"Field","name":{"kind":"Name","value":"planOutput"}}]}}]}}]} as unknown as DocumentNode<PlanExecutionRequestQuery, PlanExecutionRequestQueryVariables>;
export const CreateModulePropagationDriftCheckRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModulePropagationDriftCheckRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModulePropagationDriftCheckRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationDriftCheckRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]} as unknown as DocumentNode<CreateModulePropagationDriftCheckRequestMutation, CreateModulePropagationDriftCheckRequestMutationVariables>;
export const CreateModulePropagationExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModulePropagationExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModulePropagationExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"modulePropagationId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationId"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequestId"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]} as unknown as DocumentNode<CreateModulePropagationExecutionRequestMutation, CreateModulePropagationExecutionRequestMutationVariables>;
export const CreateTerraformDriftCheckRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createTerraformDriftCheckRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createTerraformDriftCheckRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"terraformDriftCheckRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"moduleAssignmentId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentId"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"terraformDriftCheckRequestId"}}]}}]}}]} as unknown as DocumentNode<CreateTerraformDriftCheckRequestMutation, CreateTerraformDriftCheckRequestMutationVariables>;
export const CreateTerraformExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createTerraformExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"destroy"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Boolean"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createTerraformExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"terraformExecutionRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"moduleAssignmentId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentId"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"destroy"},"value":{"kind":"Variable","name":{"kind":"Name","value":"destroy"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"terraformExecutionRequestId"}}]}}]}}]} as unknown as DocumentNode<CreateTerraformExecutionRequestMutation, CreateTerraformExecutionRequestMutationVariables>;