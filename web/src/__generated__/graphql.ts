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
  additionalArguments?: Maybe<Scalars['String']>;
  applyOutput?: Maybe<Scalars['String']>;
  completedAt?: Maybe<Scalars['Time']>;
  id: Scalars['ID'];
  initOutput?: Maybe<Scalars['String']>;
  moduleAssignment: ModuleAssignment;
  moduleAssignmentID: Scalars['ID'];
  startedAt?: Maybe<Scalars['Time']>;
  status: RequestStatus;
  terraformConfiguration: Scalars['String'];
  terraformExecutionRequestID: Scalars['ID'];
  terraformPlan: Scalars['String'];
  terraformVersion: Scalars['String'];
};

export type ApplyExecutionRequestFilters = {
  completedAfter?: InputMaybe<Scalars['Time']>;
  completedBefore?: InputMaybe<Scalars['Time']>;
  destroy?: InputMaybe<Scalars['Boolean']>;
  startedAfter?: InputMaybe<Scalars['Time']>;
  startedBefore?: InputMaybe<Scalars['Time']>;
  status?: InputMaybe<RequestStatus>;
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
  arguments: Array<Argument>;
  awsProviderConfigurations?: Maybe<Array<AwsProviderConfiguration>>;
  description: Scalars['String'];
  gcpProviderConfigurations?: Maybe<Array<GcpProviderConfiguration>>;
  id: Scalars['ID'];
  moduleGroup: ModuleGroup;
  moduleGroupID: Scalars['ID'];
  modulePropagation?: Maybe<ModulePropagation>;
  modulePropagationID?: Maybe<Scalars['ID']>;
  moduleVersion: ModuleVersion;
  moduleVersionID: Scalars['ID'];
  name: Scalars['String'];
  orgAccount: OrgAccount;
  orgAccountID: Scalars['ID'];
  remoteStateBucket: Scalars['String'];
  remoteStateKey: Scalars['String'];
  remoteStateRegion: Scalars['String'];
  status: ModuleAssignmentStatus;
  terraformConfiguration: Scalars['String'];
  terraformDriftCheckRequests: Array<TerraformDriftCheckRequest>;
  terraformExecutionRequests: Array<TerraformExecutionRequest>;
};


export type ModuleAssignmentTerraformDriftCheckRequestsArgs = {
  filters?: InputMaybe<TerraformDriftCheckRequestFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type ModuleAssignmentTerraformExecutionRequestsArgs = {
  filters?: InputMaybe<TerraformExecutionRequestFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};

export type ModuleAssignmentFilters = {
  descriptionContains?: InputMaybe<Scalars['String']>;
  isPropagated?: InputMaybe<Scalars['Boolean']>;
  nameContains?: InputMaybe<Scalars['String']>;
  status?: InputMaybe<ModuleAssignmentStatus>;
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
  moduleVersionID?: InputMaybe<Scalars['ID']>;
  name?: InputMaybe<Scalars['String']>;
};

export type ModuleGroup = {
  __typename?: 'ModuleGroup';
  cloudPlatform: CloudPlatform;
  id: Scalars['ID'];
  moduleAssignments: Array<ModuleAssignment>;
  modulePropagations: Array<ModulePropagation>;
  name: Scalars['String'];
  versions: Array<ModuleVersion>;
};


export type ModuleGroupModuleAssignmentsArgs = {
  filters?: InputMaybe<ModuleAssignmentFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type ModuleGroupModulePropagationsArgs = {
  filters?: InputMaybe<ModulePropagationFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type ModuleGroupVersionsArgs = {
  filters?: InputMaybe<ModuleVersionFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};

export type ModuleGroupFilters = {
  cloudPlatform?: InputMaybe<CloudPlatform>;
  nameContains?: InputMaybe<Scalars['String']>;
};

export type ModulePropagation = {
  __typename?: 'ModulePropagation';
  arguments: Array<Argument>;
  awsProviderConfigurations?: Maybe<Array<AwsProviderConfiguration>>;
  description: Scalars['String'];
  driftCheckRequests: Array<ModulePropagationDriftCheckRequest>;
  executionRequests: Array<ModulePropagationExecutionRequest>;
  gcpProviderConfigurations?: Maybe<Array<GcpProviderConfiguration>>;
  id: Scalars['ID'];
  moduleAssignments: Array<ModuleAssignment>;
  moduleGroup: ModuleGroup;
  moduleGroupID: Scalars['ID'];
  moduleVersion: ModuleVersion;
  moduleVersionID: Scalars['ID'];
  name: Scalars['String'];
  orgDimension: OrgDimension;
  orgDimensionID: Scalars['ID'];
  orgUnit: OrgUnit;
  orgUnitID: Scalars['ID'];
};


export type ModulePropagationDriftCheckRequestsArgs = {
  filters?: InputMaybe<ModulePropagationDriftCheckRequestFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type ModulePropagationExecutionRequestsArgs = {
  filters?: InputMaybe<ModulePropagationExecutionRequestFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type ModulePropagationModuleAssignmentsArgs = {
  filters?: InputMaybe<ModuleAssignmentFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};

export type ModulePropagationDriftCheckRequest = {
  __typename?: 'ModulePropagationDriftCheckRequest';
  completedAt?: Maybe<Scalars['Time']>;
  id: Scalars['ID'];
  modulePropagation: ModulePropagation;
  modulePropagationID: Scalars['ID'];
  startedAt?: Maybe<Scalars['Time']>;
  status: RequestStatus;
  syncStatus: TerraformDriftCheckStatus;
  terraformDriftCheckRequests: Array<TerraformDriftCheckRequest>;
};


export type ModulePropagationDriftCheckRequestTerraformDriftCheckRequestsArgs = {
  filters?: InputMaybe<TerraformDriftCheckRequestFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};

export type ModulePropagationDriftCheckRequestFilters = {
  completedAfter?: InputMaybe<Scalars['Time']>;
  completedBefore?: InputMaybe<Scalars['Time']>;
  startedAfter?: InputMaybe<Scalars['Time']>;
  startedBefore?: InputMaybe<Scalars['Time']>;
  status?: InputMaybe<RequestStatus>;
  syncStatus?: InputMaybe<TerraformDriftCheckStatus>;
};

export type ModulePropagationExecutionRequest = {
  __typename?: 'ModulePropagationExecutionRequest';
  completedAt?: Maybe<Scalars['Time']>;
  id: Scalars['ID'];
  modulePropagation: ModulePropagation;
  modulePropagationID: Scalars['ID'];
  startedAt?: Maybe<Scalars['Time']>;
  status: RequestStatus;
  terraformExecutionRequests: Array<TerraformExecutionRequest>;
};


export type ModulePropagationExecutionRequestTerraformExecutionRequestsArgs = {
  filters?: InputMaybe<TerraformExecutionRequestFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};

export type ModulePropagationExecutionRequestFilters = {
  completedAfter?: InputMaybe<Scalars['Time']>;
  completedBefore?: InputMaybe<Scalars['Time']>;
  startedAfter?: InputMaybe<Scalars['Time']>;
  startedBefore?: InputMaybe<Scalars['Time']>;
  status?: InputMaybe<RequestStatus>;
};

export type ModulePropagationFilters = {
  descriptionContains?: InputMaybe<Scalars['String']>;
  nameContains?: InputMaybe<Scalars['String']>;
};

export type ModulePropagationUpdate = {
  arguments?: InputMaybe<Array<ArgumentInput>>;
  awsProviderConfigurations?: InputMaybe<Array<AwsProviderConfigurationInput>>;
  description?: InputMaybe<Scalars['String']>;
  gcpProviderConfigurations?: InputMaybe<Array<GcpProviderConfigurationInput>>;
  moduleVersionID?: InputMaybe<Scalars['ID']>;
  name?: InputMaybe<Scalars['String']>;
  orgDimensionID?: InputMaybe<Scalars['ID']>;
  orgUnitID?: InputMaybe<Scalars['ID']>;
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
  id: Scalars['ID'];
  moduleAssignments: Array<ModuleAssignment>;
  moduleGroup: ModuleGroup;
  moduleGroupID: Scalars['ID'];
  modulePropagations: Array<ModulePropagation>;
  name: Scalars['String'];
  remoteSource: Scalars['String'];
  terraformVersion: Scalars['String'];
  variables: Array<ModuleVariable>;
};


export type ModuleVersionModuleAssignmentsArgs = {
  filters?: InputMaybe<ModuleAssignmentFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type ModuleVersionModulePropagationsArgs = {
  filters?: InputMaybe<ModulePropagationFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};

export type ModuleVersionFilters = {
  nameContains?: InputMaybe<Scalars['String']>;
  remoteSourceContains?: InputMaybe<Scalars['String']>;
  terraformVersion?: InputMaybe<Scalars['String']>;
};

export type Mutation = {
  __typename?: 'Mutation';
  addAccountToOrgUnit: Scalars['Boolean'];
  createModuleAssignment: ModuleAssignment;
  createModuleGroup: ModuleGroup;
  createModulePropagation: ModulePropagation;
  createModulePropagationDriftCheckRequest: ModulePropagationDriftCheckRequest;
  createModulePropagationExecutionRequest: ModulePropagationExecutionRequest;
  createModuleVersion: ModuleVersion;
  createOrgAccount: OrgAccount;
  createOrgDimension: OrgDimension;
  createOrgUnit: OrgUnit;
  createTerraformDriftCheckRequest: TerraformDriftCheckRequest;
  createTerraformExecutionRequest: TerraformExecutionRequest;
  deleteModuleAssignment: Scalars['Boolean'];
  deleteModuleGroup: Scalars['Boolean'];
  deleteModulePropagation: Scalars['Boolean'];
  deleteModuleVersion: Scalars['Boolean'];
  deleteOrgAccount: Scalars['Boolean'];
  deleteOrgDimension: Scalars['Boolean'];
  deleteOrgUnit: Scalars['Boolean'];
  removeAccountFromOrgUnit: Scalars['Boolean'];
  updateModuleAssignment: ModuleAssignment;
  updateModulePropagation: ModulePropagation;
  updateOrgAccount: OrgAccount;
  updateOrgUnit: OrgUnit;
};


export type MutationAddAccountToOrgUnitArgs = {
  orgAccountID: Scalars['ID'];
  orgUnitID: Scalars['ID'];
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


export type MutationCreateOrgAccountArgs = {
  orgAccount: NewOrgAccount;
};


export type MutationCreateOrgDimensionArgs = {
  orgDimension: NewOrgDimension;
};


export type MutationCreateOrgUnitArgs = {
  orgUnit: NewOrgUnit;
};


export type MutationCreateTerraformDriftCheckRequestArgs = {
  terraformDriftCheckRequest: NewTerraformDriftCheckRequest;
};


export type MutationCreateTerraformExecutionRequestArgs = {
  terraformExecutionRequest: NewTerraformExecutionRequest;
};


export type MutationDeleteModuleAssignmentArgs = {
  moduleAssignmentID: Scalars['ID'];
};


export type MutationDeleteModuleGroupArgs = {
  moduleGroupID: Scalars['ID'];
};


export type MutationDeleteModulePropagationArgs = {
  modulePropagationID: Scalars['ID'];
};


export type MutationDeleteModuleVersionArgs = {
  moduleVersionID: Scalars['ID'];
};


export type MutationDeleteOrgAccountArgs = {
  orgAccountID: Scalars['ID'];
};


export type MutationDeleteOrgDimensionArgs = {
  orgDimensionID: Scalars['ID'];
};


export type MutationDeleteOrgUnitArgs = {
  orgUnitID: Scalars['ID'];
};


export type MutationRemoveAccountFromOrgUnitArgs = {
  orgAccountID: Scalars['ID'];
  orgUnitID: Scalars['ID'];
};


export type MutationUpdateModuleAssignmentArgs = {
  moduleAssignmentID: Scalars['ID'];
  moduleAssignmentUpdate: ModuleAssignmentUpdate;
};


export type MutationUpdateModulePropagationArgs = {
  modulePropagationID: Scalars['ID'];
  update: ModulePropagationUpdate;
};


export type MutationUpdateOrgAccountArgs = {
  orgAccount: OrgAccountUpdate;
  orgAccountID: Scalars['ID'];
};


export type MutationUpdateOrgUnitArgs = {
  orgUnitID: Scalars['ID'];
  update: OrgUnitUpdate;
};

export type NewModuleAssignment = {
  arguments: Array<ArgumentInput>;
  awsProviderConfigurations?: InputMaybe<Array<AwsProviderConfigurationInput>>;
  description: Scalars['String'];
  gcpProviderConfigurations?: InputMaybe<Array<GcpProviderConfigurationInput>>;
  moduleGroupID: Scalars['ID'];
  moduleVersionID: Scalars['ID'];
  name: Scalars['String'];
  orgAccountID: Scalars['ID'];
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
  moduleGroupID: Scalars['ID'];
  moduleVersionID: Scalars['ID'];
  name: Scalars['String'];
  orgDimensionID: Scalars['ID'];
  orgUnitID: Scalars['ID'];
};

export type NewModulePropagationDriftCheckRequest = {
  modulePropagationID: Scalars['ID'];
};

export type NewModulePropagationExecutionRequest = {
  modulePropagationID: Scalars['ID'];
};

export type NewModuleVersion = {
  moduleGroupID: Scalars['ID'];
  name: Scalars['String'];
  remoteSource: Scalars['String'];
  terraformVersion: Scalars['String'];
};

export type NewOrgAccount = {
  assumeRoleName: Scalars['String'];
  cloudIdentifier: Scalars['String'];
  cloudPlatform: CloudPlatform;
  metadata: Array<MetadataInput>;
  name: Scalars['String'];
};

export type NewOrgDimension = {
  name: Scalars['String'];
};

export type NewOrgUnit = {
  name: Scalars['String'];
  orgDimensionID: Scalars['ID'];
  parentOrgUnitID: Scalars['ID'];
};

export type NewTerraformDriftCheckRequest = {
  moduleAssignmentID: Scalars['ID'];
};

export type NewTerraformExecutionRequest = {
  destroy: Scalars['Boolean'];
  moduleAssignmentID: Scalars['ID'];
};

export type OrgAccount = {
  __typename?: 'OrgAccount';
  assumeRoleName: Scalars['String'];
  cloudIdentifier: Scalars['String'];
  cloudPlatform: CloudPlatform;
  id: Scalars['ID'];
  metadata: Array<Metadata>;
  moduleAssignments: Array<ModuleAssignment>;
  name: Scalars['String'];
  orgUnits: Array<OrgUnit>;
};


export type OrgAccountModuleAssignmentsArgs = {
  filters?: InputMaybe<ModuleAssignmentFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type OrgAccountOrgUnitsArgs = {
  filters?: InputMaybe<OrgUnitFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};

export type OrgAccountFilters = {
  cloudIdentifier?: InputMaybe<Scalars['String']>;
  cloudPlatform?: InputMaybe<CloudPlatform>;
  nameContains?: InputMaybe<Scalars['String']>;
};

export type OrgAccountUpdate = {
  assumeRoleName?: InputMaybe<Scalars['String']>;
  cloudIdentifier?: InputMaybe<Scalars['String']>;
  cloudPlatform?: InputMaybe<CloudPlatform>;
  metadata?: InputMaybe<Array<MetadataInput>>;
  name?: InputMaybe<Scalars['String']>;
};

export type OrgDimension = {
  __typename?: 'OrgDimension';
  id: Scalars['ID'];
  modulePropagations: Array<ModulePropagation>;
  name: Scalars['String'];
  orgUnits: Array<OrgUnit>;
};


export type OrgDimensionModulePropagationsArgs = {
  filters?: InputMaybe<ModulePropagationFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type OrgDimensionOrgUnitsArgs = {
  filters?: InputMaybe<OrgUnitFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};

export type OrgDimensionFilters = {
  nameContains?: InputMaybe<Scalars['String']>;
};

export type OrgUnit = {
  __typename?: 'OrgUnit';
  children: Array<OrgUnit>;
  downstreamOrgUnits: Array<OrgUnit>;
  hierarchy: Scalars['String'];
  id: Scalars['ID'];
  modulePropagations: Array<ModulePropagation>;
  name: Scalars['String'];
  orgAccounts: Array<OrgAccount>;
  orgDimension: OrgDimension;
  orgDimensionID: Scalars['ID'];
  parentOrgUnit?: Maybe<OrgUnit>;
  parentOrgUnitID?: Maybe<Scalars['ID']>;
  upstreamOrgUnits: Array<OrgUnit>;
};


export type OrgUnitChildrenArgs = {
  filters?: InputMaybe<OrgUnitFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type OrgUnitDownstreamOrgUnitsArgs = {
  filters?: InputMaybe<OrgUnitFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type OrgUnitModulePropagationsArgs = {
  filters?: InputMaybe<ModulePropagationFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type OrgUnitOrgAccountsArgs = {
  filters?: InputMaybe<OrgAccountFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type OrgUnitUpstreamOrgUnitsArgs = {
  filters?: InputMaybe<OrgUnitFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};

export type OrgUnitFilters = {
  nameContains?: InputMaybe<Scalars['String']>;
};

export type OrgUnitUpdate = {
  Name?: InputMaybe<Scalars['String']>;
  ParentOrgUnitID?: InputMaybe<Scalars['ID']>;
};

export type PlanExecutionRequest = {
  __typename?: 'PlanExecutionRequest';
  additionalArguments?: Maybe<Scalars['String']>;
  completedAt?: Maybe<Scalars['Time']>;
  id: Scalars['ID'];
  initOutput?: Maybe<Scalars['String']>;
  moduleAssignment: ModuleAssignment;
  moduleAssignmentID: Scalars['ID'];
  planFile?: Maybe<Scalars['String']>;
  planJSON?: Maybe<Scalars['String']>;
  planOutput?: Maybe<Scalars['String']>;
  startedAt?: Maybe<Scalars['Time']>;
  status: RequestStatus;
  terraformConfiguration: Scalars['String'];
  terraformDriftCheckRequestID?: Maybe<Scalars['ID']>;
  terraformExecutionRequestID?: Maybe<Scalars['ID']>;
  terraformVersion: Scalars['String'];
};

export type PlanExecutionRequestFilters = {
  completedAfter?: InputMaybe<Scalars['Time']>;
  completedBefore?: InputMaybe<Scalars['Time']>;
  destroy?: InputMaybe<Scalars['Boolean']>;
  startedAfter?: InputMaybe<Scalars['Time']>;
  startedBefore?: InputMaybe<Scalars['Time']>;
  status?: InputMaybe<RequestStatus>;
};

export type Query = {
  __typename?: 'Query';
  applyExecutionRequest: ApplyExecutionRequest;
  moduleAssignment: ModuleAssignment;
  moduleAssignments: Array<ModuleAssignment>;
  moduleGroup: ModuleGroup;
  moduleGroups: Array<ModuleGroup>;
  modulePropagation: ModulePropagation;
  modulePropagationDriftCheckRequest: ModulePropagationDriftCheckRequest;
  modulePropagationDriftCheckRequests: Array<ModulePropagationDriftCheckRequest>;
  modulePropagationExecutionRequest: ModulePropagationExecutionRequest;
  modulePropagationExecutionRequests: Array<ModulePropagationExecutionRequest>;
  modulePropagations: Array<ModulePropagation>;
  moduleVersion: ModuleVersion;
  moduleVersions: Array<ModuleVersion>;
  orgAccount: OrgAccount;
  orgAccounts: Array<OrgAccount>;
  orgDimension: OrgDimension;
  orgDimensions: Array<OrgDimension>;
  orgUnit: OrgUnit;
  planExecutionRequest: PlanExecutionRequest;
};


export type QueryApplyExecutionRequestArgs = {
  applyExecutionRequestID: Scalars['ID'];
};


export type QueryModuleAssignmentArgs = {
  moduleAssignmentID: Scalars['ID'];
};


export type QueryModuleAssignmentsArgs = {
  filters?: InputMaybe<ModuleAssignmentFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type QueryModuleGroupArgs = {
  moduleGroupID: Scalars['ID'];
};


export type QueryModuleGroupsArgs = {
  filters?: InputMaybe<ModuleGroupFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type QueryModulePropagationArgs = {
  modulePropagationID: Scalars['ID'];
};


export type QueryModulePropagationDriftCheckRequestArgs = {
  modulePropagationDriftCheckRequestID: Scalars['ID'];
};


export type QueryModulePropagationDriftCheckRequestsArgs = {
  filters?: InputMaybe<ModulePropagationDriftCheckRequestFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type QueryModulePropagationExecutionRequestArgs = {
  modulePropagationExecutionRequestID: Scalars['ID'];
};


export type QueryModulePropagationExecutionRequestsArgs = {
  filters?: InputMaybe<ModulePropagationExecutionRequestFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type QueryModulePropagationsArgs = {
  filters?: InputMaybe<ModulePropagationFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type QueryModuleVersionArgs = {
  moduleVersionID: Scalars['ID'];
};


export type QueryModuleVersionsArgs = {
  filters?: InputMaybe<ModuleVersionFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type QueryOrgAccountArgs = {
  orgAccountID: Scalars['ID'];
};


export type QueryOrgAccountsArgs = {
  filters?: InputMaybe<OrgAccountFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type QueryOrgDimensionArgs = {
  orgDimensionID: Scalars['ID'];
};


export type QueryOrgDimensionsArgs = {
  filters?: InputMaybe<OrgDimensionFilters>;
  limit?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};


export type QueryOrgUnitArgs = {
  orgUnitID: Scalars['ID'];
};


export type QueryPlanExecutionRequestArgs = {
  planExecutionRequestID: Scalars['ID'];
};

export enum RequestStatus {
  Failed = 'FAILED',
  Pending = 'PENDING',
  Running = 'RUNNING',
  Succeeded = 'SUCCEEDED'
}

export type TerraformDriftCheckRequest = {
  __typename?: 'TerraformDriftCheckRequest';
  completedAt?: Maybe<Scalars['Time']>;
  destroy: Scalars['Boolean'];
  id: Scalars['ID'];
  moduleAssignment: ModuleAssignment;
  moduleAssignmentID: Scalars['ID'];
  modulePropagation?: Maybe<ModulePropagation>;
  modulePropagationDriftCheckRequest?: Maybe<ModulePropagationDriftCheckRequest>;
  modulePropagationDriftCheckRequestID?: Maybe<Scalars['ID']>;
  modulePropagationID?: Maybe<Scalars['ID']>;
  planExecutionRequest?: Maybe<PlanExecutionRequest>;
  startedAt?: Maybe<Scalars['Time']>;
  status: RequestStatus;
  syncStatus: TerraformDriftCheckStatus;
};

export type TerraformDriftCheckRequestFilters = {
  completedAfter?: InputMaybe<Scalars['Time']>;
  completedBefore?: InputMaybe<Scalars['Time']>;
  destroy?: InputMaybe<Scalars['Boolean']>;
  startedAfter?: InputMaybe<Scalars['Time']>;
  startedBefore?: InputMaybe<Scalars['Time']>;
  status?: InputMaybe<RequestStatus>;
  syncStatus?: InputMaybe<TerraformDriftCheckStatus>;
};

export enum TerraformDriftCheckStatus {
  InSync = 'IN_SYNC',
  OutOfSync = 'OUT_OF_SYNC',
  Pending = 'PENDING'
}

export type TerraformExecutionRequest = {
  __typename?: 'TerraformExecutionRequest';
  applyExecutionRequest?: Maybe<ApplyExecutionRequest>;
  completedAt?: Maybe<Scalars['Time']>;
  destroy: Scalars['Boolean'];
  id: Scalars['ID'];
  moduleAssignment: ModuleAssignment;
  moduleAssignmentID: Scalars['ID'];
  modulePropagation?: Maybe<ModulePropagation>;
  modulePropagationExecutionRequest?: Maybe<ModulePropagationExecutionRequest>;
  modulePropagationExecutionRequestID?: Maybe<Scalars['ID']>;
  modulePropagationID?: Maybe<Scalars['ID']>;
  planExecutionRequest?: Maybe<PlanExecutionRequest>;
  startedAt?: Maybe<Scalars['Time']>;
  status: RequestStatus;
};

export type TerraformExecutionRequestFilters = {
  completedAfter?: InputMaybe<Scalars['Time']>;
  completedBefore?: InputMaybe<Scalars['Time']>;
  destroy?: InputMaybe<Scalars['Boolean']>;
  startedAfter?: InputMaybe<Scalars['Time']>;
  startedBefore?: InputMaybe<Scalars['Time']>;
  status?: InputMaybe<RequestStatus>;
};

export type ApplyExecutionRequestQueryVariables = Exact<{
  applyExecutionRequestID: Scalars['ID'];
}>;


export type ApplyExecutionRequestQuery = { __typename?: 'Query', applyExecutionRequest: { __typename?: 'ApplyExecutionRequest', id: string, status: RequestStatus, startedAt?: any | null, completedAt?: any | null, terraformConfiguration: string, initOutput?: string | null, applyOutput?: string | null, moduleAssignment: { __typename?: 'ModuleAssignment', name: string, modulePropagation?: { __typename?: 'ModulePropagation', name: string } | null, orgAccount: { __typename?: 'OrgAccount', id: string, name: string, cloudPlatform: CloudPlatform }, moduleGroup: { __typename?: 'ModuleGroup', id: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string } } } };

export type RemoveAccountFromOrgUnitMutationVariables = Exact<{
  orgUnitID: Scalars['ID'];
  orgAccountID: Scalars['ID'];
}>;


export type RemoveAccountFromOrgUnitMutation = { __typename?: 'Mutation', removeAccountFromOrgUnit: boolean };

export type ModuleAssignmentQueryVariables = Exact<{
  moduleAssignmentID: Scalars['ID'];
}>;


export type ModuleAssignmentQuery = { __typename?: 'Query', moduleAssignment: { __typename?: 'ModuleAssignment', id: string, name: string, status: ModuleAssignmentStatus, terraformConfiguration: string, modulePropagation?: { __typename?: 'ModulePropagation', id: string, name: string } | null, orgAccount: { __typename?: 'OrgAccount', id: string, name: string, cloudPlatform: CloudPlatform }, moduleGroup: { __typename?: 'ModuleGroup', id: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string }, terraformExecutionRequests: Array<{ __typename?: 'TerraformExecutionRequest', id: string, status: RequestStatus, startedAt?: any | null, destroy: boolean, moduleAssignment: { __typename?: 'ModuleAssignment', id: string, orgAccount: { __typename?: 'OrgAccount', id: string, cloudPlatform: CloudPlatform, name: string } }, planExecutionRequest?: { __typename?: 'PlanExecutionRequest', id: string, status: RequestStatus, startedAt?: any | null } | null, applyExecutionRequest?: { __typename?: 'ApplyExecutionRequest', id: string, status: RequestStatus, startedAt?: any | null } | null }>, terraformDriftCheckRequests: Array<{ __typename?: 'TerraformDriftCheckRequest', id: string, status: RequestStatus, startedAt?: any | null, destroy: boolean, syncStatus: TerraformDriftCheckStatus, moduleAssignment: { __typename?: 'ModuleAssignment', id: string, orgAccount: { __typename?: 'OrgAccount', id: string, cloudPlatform: CloudPlatform, name: string } }, planExecutionRequest?: { __typename?: 'PlanExecutionRequest', id: string, status: RequestStatus, startedAt?: any | null } | null }> } };

export type ModuleAssignmentsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModuleAssignmentsQuery = { __typename?: 'Query', unpropagated: Array<{ __typename?: 'ModuleAssignment', id: string, name: string, description: string, moduleGroup: { __typename?: 'ModuleGroup', id: string, name: string, cloudPlatform: CloudPlatform }, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string }, orgAccount: { __typename?: 'OrgAccount', id: string, name: string, cloudPlatform: CloudPlatform } }> };

export type ModuleGroupsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModuleGroupsQuery = { __typename?: 'Query', moduleGroups: Array<{ __typename?: 'ModuleGroup', id: string, cloudPlatform: CloudPlatform, name: string, versions: Array<{ __typename?: 'ModuleVersion', id: string, remoteSource: string, terraformVersion: string, name: string }> }> };

export type ModuleGroupQueryVariables = Exact<{
  moduleGroupID: Scalars['ID'];
}>;


export type ModuleGroupQuery = { __typename?: 'Query', moduleGroup: { __typename?: 'ModuleGroup', id: string, name: string, cloudPlatform: CloudPlatform, versions: Array<{ __typename?: 'ModuleVersion', id: string, name: string, remoteSource: string, terraformVersion: string }>, modulePropagations: Array<{ __typename?: 'ModulePropagation', name: string, id: string, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string }, orgUnit: { __typename?: 'OrgUnit', id: string, name: string }, orgDimension: { __typename?: 'OrgDimension', id: string, name: string } }>, moduleAssignments: Array<{ __typename?: 'ModuleAssignment', id: string, status: ModuleAssignmentStatus, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string }, modulePropagation?: { __typename?: 'ModulePropagation', id: string, name: string } | null, orgAccount: { __typename?: 'OrgAccount', id: string, name: string } }> } };

export type ModulePropagationDriftCheckRequestQueryVariables = Exact<{
  modulePropagationDriftCheckRequestID: Scalars['ID'];
}>;


export type ModulePropagationDriftCheckRequestQuery = { __typename?: 'Query', modulePropagationDriftCheckRequest: { __typename?: 'ModulePropagationDriftCheckRequest', id: string, startedAt?: any | null, status: RequestStatus, terraformDriftCheckRequests: Array<{ __typename?: 'TerraformDriftCheckRequest', id: string, status: RequestStatus, startedAt?: any | null, destroy: boolean, syncStatus: TerraformDriftCheckStatus, moduleAssignment: { __typename?: 'ModuleAssignment', id: string, orgAccount: { __typename?: 'OrgAccount', id: string, cloudPlatform: CloudPlatform, name: string } }, planExecutionRequest?: { __typename?: 'PlanExecutionRequest', id: string, status: RequestStatus, startedAt?: any | null } | null }> } };

export type ModulePropagationExecutionRequestQueryVariables = Exact<{
  modulePropagationExecutionRequestID: Scalars['ID'];
}>;


export type ModulePropagationExecutionRequestQuery = { __typename?: 'Query', modulePropagationExecutionRequest: { __typename?: 'ModulePropagationExecutionRequest', id: string, startedAt?: any | null, status: RequestStatus, modulePropagation: { __typename?: 'ModulePropagation', name: string }, terraformExecutionRequests: Array<{ __typename?: 'TerraformExecutionRequest', id: string, status: RequestStatus, startedAt?: any | null, destroy: boolean, moduleAssignment: { __typename?: 'ModuleAssignment', id: string, orgAccount: { __typename?: 'OrgAccount', id: string, cloudPlatform: CloudPlatform, name: string } }, planExecutionRequest?: { __typename?: 'PlanExecutionRequest', id: string, status: RequestStatus, startedAt?: any | null } | null, applyExecutionRequest?: { __typename?: 'ApplyExecutionRequest', id: string, status: RequestStatus, startedAt?: any | null } | null }> } };

export type ModulePropagationQueryVariables = Exact<{
  modulePropagationID: Scalars['ID'];
}>;


export type ModulePropagationQuery = { __typename?: 'Query', modulePropagation: { __typename?: 'ModulePropagation', id: string, name: string, description: string, moduleGroup: { __typename?: 'ModuleGroup', id: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string }, orgUnit: { __typename?: 'OrgUnit', id: string, name: string, orgDimension: { __typename?: 'OrgDimension', id: string, name: string }, downstreamOrgUnits: Array<{ __typename?: 'OrgUnit', id: string, name: string, orgDimension: { __typename?: 'OrgDimension', id: string, name: string } }> }, executionRequests: Array<{ __typename?: 'ModulePropagationExecutionRequest', id: string, modulePropagationID: string, startedAt?: any | null, completedAt?: any | null, status: RequestStatus }>, driftCheckRequests: Array<{ __typename?: 'ModulePropagationDriftCheckRequest', id: string, modulePropagationID: string, startedAt?: any | null, status: RequestStatus, syncStatus: TerraformDriftCheckStatus }>, moduleAssignments: Array<{ __typename?: 'ModuleAssignment', id: string, modulePropagationID?: string | null, status: ModuleAssignmentStatus, orgAccount: { __typename?: 'OrgAccount', id: string, name: string } }> } };

export type ModulePropagationUpdateOptionsQueryVariables = Exact<{
  moduleGroupID: Scalars['ID'];
}>;


export type ModulePropagationUpdateOptionsQuery = { __typename?: 'Query', orgDimensions: Array<{ __typename?: 'OrgDimension', id: string, name: string, orgUnits: Array<{ __typename?: 'OrgUnit', id: string, name: string }> }>, moduleGroup: { __typename?: 'ModuleGroup', versions: Array<{ __typename?: 'ModuleVersion', id: string, name: string }> } };

export type UpdateModulePropagationMutationVariables = Exact<{
  modulePropagationID: Scalars['ID'];
  update: ModulePropagationUpdate;
}>;


export type UpdateModulePropagationMutation = { __typename?: 'Mutation', updateModulePropagation: { __typename?: 'ModulePropagation', id: string } };

export type ModulePropagationsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModulePropagationsQuery = { __typename?: 'Query', modulePropagations: Array<{ __typename?: 'ModulePropagation', id: string, name: string, moduleGroup: { __typename?: 'ModuleGroup', id: string, name: string, cloudPlatform: CloudPlatform }, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string }, orgUnit: { __typename?: 'OrgUnit', id: string, name: string }, orgDimension: { __typename?: 'OrgDimension', id: string, name: string } }> };

export type ModuleVersionQueryVariables = Exact<{
  moduleVersionID: Scalars['ID'];
}>;


export type ModuleVersionQuery = { __typename?: 'Query', moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string, remoteSource: string, terraformVersion: string, moduleGroup: { __typename?: 'ModuleGroup', id: string, cloudPlatform: CloudPlatform, name: string }, variables: Array<{ __typename?: 'ModuleVariable', name: string, type: string, description: string, default?: string | null }>, modulePropagations: Array<{ __typename?: 'ModulePropagation', name: string, description: string, id: string, orgUnit: { __typename?: 'OrgUnit', id: string, name: string }, orgDimension: { __typename?: 'OrgDimension', id: string, name: string } }>, moduleAssignments: Array<{ __typename?: 'ModuleAssignment', id: string, name: string, description: string, orgAccount: { __typename?: 'OrgAccount', id: string, name: string, cloudPlatform: CloudPlatform } }> } };

export type CreateModuleAssignmentMutationVariables = Exact<{
  moduleAssignment: NewModuleAssignment;
}>;


export type CreateModuleAssignmentMutation = { __typename?: 'Mutation', createModuleAssignment: { __typename?: 'ModuleAssignment', id: string } };

export type ModuleAssignmentOptionsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModuleAssignmentOptionsQuery = { __typename?: 'Query', orgAccounts: Array<{ __typename?: 'OrgAccount', id: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string }>, moduleGroups: Array<{ __typename?: 'ModuleGroup', id: string, name: string, cloudPlatform: CloudPlatform, versions: Array<{ __typename?: 'ModuleVersion', id: string, name: string, variables: Array<{ __typename?: 'ModuleVariable', name: string, type: string, default?: string | null, description: string }> }> }> };

export type CreateModuleGroupMutationVariables = Exact<{
  moduleGroup: NewModuleGroup;
}>;


export type CreateModuleGroupMutation = { __typename?: 'Mutation', createModuleGroup: { __typename?: 'ModuleGroup', id: string } };

export type CreateModulePropagationMutationVariables = Exact<{
  modulePropagation: NewModulePropagation;
}>;


export type CreateModulePropagationMutation = { __typename?: 'Mutation', createModulePropagation: { __typename?: 'ModulePropagation', id: string } };

export type ModulePropagationOptionsQueryVariables = Exact<{ [key: string]: never; }>;


export type ModulePropagationOptionsQuery = { __typename?: 'Query', orgDimensions: Array<{ __typename?: 'OrgDimension', id: string, name: string, orgUnits: Array<{ __typename?: 'OrgUnit', id: string, name: string }> }>, moduleGroups: Array<{ __typename?: 'ModuleGroup', id: string, name: string, cloudPlatform: CloudPlatform, versions: Array<{ __typename?: 'ModuleVersion', id: string, name: string, variables: Array<{ __typename?: 'ModuleVariable', name: string, type: string, default?: string | null, description: string }> }> }> };

export type CreateModuleVersionMutationVariables = Exact<{
  moduleVersion: NewModuleVersion;
}>;


export type CreateModuleVersionMutation = { __typename?: 'Mutation', createModuleVersion: { __typename?: 'ModuleVersion', id: string } };

export type CreateOrgAccountMutationVariables = Exact<{
  orgAccount: NewOrgAccount;
}>;


export type CreateOrgAccountMutation = { __typename?: 'Mutation', createOrgAccount: { __typename?: 'OrgAccount', id: string } };

export type CreateOrgDimensionMutationVariables = Exact<{
  orgDimension: NewOrgDimension;
}>;


export type CreateOrgDimensionMutation = { __typename?: 'Mutation', createOrgDimension: { __typename?: 'OrgDimension', id: string } };

export type CreateOrgUnitMutationVariables = Exact<{
  orgUnit: NewOrgUnit;
}>;


export type CreateOrgUnitMutation = { __typename?: 'Mutation', createOrgUnit: { __typename?: 'OrgUnit', id: string } };

export type AddAccountToOrgUnitMutationVariables = Exact<{
  orgUnitID: Scalars['ID'];
  orgAccountID: Scalars['ID'];
}>;


export type AddAccountToOrgUnitMutation = { __typename?: 'Mutation', addAccountToOrgUnit: boolean };

export type OrgDimensionsAndUnitsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrgDimensionsAndUnitsQuery = { __typename?: 'Query', orgDimensions: Array<{ __typename?: 'OrgDimension', id: string, name: string, orgUnits: Array<{ __typename?: 'OrgUnit', id: string, name: string }> }> };

export type OrgAccountsAndMembershipsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrgAccountsAndMembershipsQuery = { __typename?: 'Query', orgAccounts: Array<{ __typename?: 'OrgAccount', id: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string, orgUnits: Array<{ __typename?: 'OrgUnit', id: string, orgDimensionID: string }> }> };

export type OrgAccountQueryVariables = Exact<{
  orgAccountID: Scalars['ID'];
}>;


export type OrgAccountQuery = { __typename?: 'Query', orgAccount: { __typename?: 'OrgAccount', id: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string, orgUnits: Array<{ __typename?: 'OrgUnit', id: string, name: string, orgDimension: { __typename?: 'OrgDimension', id: string, name: string } }>, moduleAssignments: Array<{ __typename?: 'ModuleAssignment', id: string, name: string, status: ModuleAssignmentStatus, moduleGroup: { __typename?: 'ModuleGroup', id: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string }, modulePropagation?: { __typename?: 'ModulePropagation', id: string, name: string, orgUnit: { __typename?: 'OrgUnit', id: string, name: string }, orgDimension: { __typename?: 'OrgDimension', id: string, name: string } } | null }> } };

export type OrgAccountsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrgAccountsQuery = { __typename?: 'Query', orgAccounts: Array<{ __typename?: 'OrgAccount', id: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string }> };

export type OrgDimensionQueryVariables = Exact<{
  orgDimensionID: Scalars['ID'];
}>;


export type OrgDimensionQuery = { __typename?: 'Query', orgDimension: { __typename?: 'OrgDimension', id: string, name: string, orgUnits: Array<{ __typename?: 'OrgUnit', id: string, name: string, parentOrgUnitID?: string | null, hierarchy: string }> } };

export type OrgDimensionsQueryVariables = Exact<{ [key: string]: never; }>;


export type OrgDimensionsQuery = { __typename?: 'Query', orgDimensions: Array<{ __typename?: 'OrgDimension', id: string, name: string }> };

export type OrgUnitQueryVariables = Exact<{
  orgUnitID: Scalars['ID'];
}>;


export type OrgUnitQuery = { __typename?: 'Query', orgUnit: { __typename?: 'OrgUnit', id: string, name: string, orgDimension: { __typename?: 'OrgDimension', id: string, name: string }, upstreamOrgUnits: Array<{ __typename?: 'OrgUnit', id: string, name: string, modulePropagations: Array<{ __typename?: 'ModulePropagation', id: string, name: string, description: string, moduleGroup: { __typename?: 'ModuleGroup', id: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string } }> }>, orgAccounts: Array<{ __typename?: 'OrgAccount', id: string, name: string, cloudPlatform: CloudPlatform, cloudIdentifier: string }>, modulePropagations: Array<{ __typename?: 'ModulePropagation', id: string, name: string, description: string, moduleGroup: { __typename?: 'ModuleGroup', id: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string } }> } };

export type PlanExecutionRequestQueryVariables = Exact<{
  planExecutionRequestID: Scalars['ID'];
}>;


export type PlanExecutionRequestQuery = { __typename?: 'Query', planExecutionRequest: { __typename?: 'PlanExecutionRequest', id: string, status: RequestStatus, startedAt?: any | null, terraformConfiguration: string, initOutput?: string | null, planOutput?: string | null, moduleAssignment: { __typename?: 'ModuleAssignment', id: string, name: string, modulePropagation?: { __typename?: 'ModulePropagation', id: string, name: string } | null, orgAccount: { __typename?: 'OrgAccount', id: string, name: string, cloudPlatform: CloudPlatform }, moduleGroup: { __typename?: 'ModuleGroup', id: string, name: string }, moduleVersion: { __typename?: 'ModuleVersion', id: string, name: string } } } };

export type CreateModulePropagationDriftCheckRequestMutationVariables = Exact<{
  modulePropagationID: Scalars['ID'];
}>;


export type CreateModulePropagationDriftCheckRequestMutation = { __typename?: 'Mutation', createModulePropagationDriftCheckRequest: { __typename?: 'ModulePropagationDriftCheckRequest', id: string, status: RequestStatus } };

export type CreateModulePropagationExecutionRequestMutationVariables = Exact<{
  modulePropagationID: Scalars['ID'];
}>;


export type CreateModulePropagationExecutionRequestMutation = { __typename?: 'Mutation', createModulePropagationExecutionRequest: { __typename?: 'ModulePropagationExecutionRequest', id: string, status: RequestStatus } };

export type CreateTerraformDriftCheckRequestMutationVariables = Exact<{
  moduleAssignmentID: Scalars['ID'];
}>;


export type CreateTerraformDriftCheckRequestMutation = { __typename?: 'Mutation', createTerraformDriftCheckRequest: { __typename?: 'TerraformDriftCheckRequest', id: string } };

export type CreateTerraformExecutionRequestMutationVariables = Exact<{
  moduleAssignmentID: Scalars['ID'];
  destroy: Scalars['Boolean'];
}>;


export type CreateTerraformExecutionRequestMutation = { __typename?: 'Mutation', createTerraformExecutionRequest: { __typename?: 'TerraformExecutionRequest', id: string } };


export const ApplyExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"applyExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"applyExecutionRequestID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"applyExecutionRequestID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"applyExecutionRequestID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}},{"kind":"Field","name":{"kind":"Name","value":"completedAt"}},{"kind":"Field","name":{"kind":"Name","value":"terraformConfiguration"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"initOutput"}},{"kind":"Field","name":{"kind":"Name","value":"applyOutput"}}]}}]}}]} as unknown as DocumentNode<ApplyExecutionRequestQuery, ApplyExecutionRequestQueryVariables>;
export const RemoveAccountFromOrgUnitDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"removeAccountFromOrgUnit"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"removeAccountFromOrgUnit"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgUnitID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitID"}}},{"kind":"Argument","name":{"kind":"Name","value":"orgAccountID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountID"}}}]}]}}]} as unknown as DocumentNode<RemoveAccountFromOrgUnitMutation, RemoveAccountFromOrgUnitMutationVariables>;
export const ModuleAssignmentDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleAssignment"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleAssignmentID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"terraformConfiguration"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"terraformExecutionRequests"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"5"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}},{"kind":"Field","name":{"kind":"Name","value":"destroy"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"terraformDriftCheckRequests"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"5"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}},{"kind":"Field","name":{"kind":"Name","value":"destroy"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"syncStatus"}}]}}]}}]}}]} as unknown as DocumentNode<ModuleAssignmentQuery, ModuleAssignmentQueryVariables>;
export const ModuleAssignmentsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleAssignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","alias":{"kind":"Name","value":"unpropagated"},"name":{"kind":"Name","value":"moduleAssignments"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"filters"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"isPropagated"},"value":{"kind":"BooleanValue","value":false}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}}]}}]}}]} as unknown as DocumentNode<ModuleAssignmentsQuery, ModuleAssignmentsQueryVariables>;
export const ModuleGroupsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroups"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"remoteSource"}},{"kind":"Field","name":{"kind":"Name","value":"terraformVersion"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<ModuleGroupsQuery, ModuleGroupsQueryVariables>;
export const ModuleGroupDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleGroup"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleGroupID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"remoteSource"}},{"kind":"Field","name":{"kind":"Name","value":"terraformVersion"}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]}}]} as unknown as DocumentNode<ModuleGroupQuery, ModuleGroupQueryVariables>;
export const ModulePropagationDriftCheckRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationDriftCheckRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationDriftCheckRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationDriftCheckRequestID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"terraformDriftCheckRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}},{"kind":"Field","name":{"kind":"Name","value":"destroy"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"syncStatus"}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationDriftCheckRequestQuery, ModulePropagationDriftCheckRequestQueryVariables>;
export const ModulePropagationExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationExecutionRequestID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationExecutionRequestID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationExecutionRequestID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"terraformExecutionRequests"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}},{"kind":"Field","name":{"kind":"Name","value":"destroy"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"applyExecutionRequest"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationExecutionRequestQuery, ModulePropagationExecutionRequestQueryVariables>;
export const ModulePropagationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"downstreamOrgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"executionRequests"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"5"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationID"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}},{"kind":"Field","name":{"kind":"Name","value":"completedAt"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}},{"kind":"Field","name":{"kind":"Name","value":"driftCheckRequests"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"5"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationID"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"syncStatus"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagationID"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationQuery, ModulePropagationQueryVariables>;
export const ModulePropagationUpdateOptionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationUpdateOptions"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleGroupID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroupID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationUpdateOptionsQuery, ModulePropagationUpdateOptionsQueryVariables>;
export const UpdateModulePropagationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateModulePropagation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"update"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ModulePropagationUpdate"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateModulePropagation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationID"}}},{"kind":"Argument","name":{"kind":"Name","value":"update"},"value":{"kind":"Variable","name":{"kind":"Name","value":"update"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<UpdateModulePropagationMutation, UpdateModulePropagationMutationVariables>;
export const ModulePropagationsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationsQuery, ModulePropagationsQueryVariables>;
export const ModuleVersionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleVersion"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleVersionID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleVersionID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleVersionID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"remoteSource"}},{"kind":"Field","name":{"kind":"Name","value":"terraformVersion"}},{"kind":"Field","name":{"kind":"Name","value":"variables"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"default"}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleVersionQuery, ModuleVersionQueryVariables>;
export const CreateModuleAssignmentDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModuleAssignment"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignment"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewModuleAssignment"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModuleAssignment"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleAssignment"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignment"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateModuleAssignmentMutation, CreateModuleAssignmentMutationVariables>;
export const ModuleAssignmentOptionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"moduleAssignmentOptions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"variables"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"default"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModuleAssignmentOptionsQuery, ModuleAssignmentOptionsQueryVariables>;
export const CreateModuleGroupDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModuleGroup"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroup"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewModuleGroup"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModuleGroup"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleGroup"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleGroup"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateModuleGroupMutation, CreateModuleGroupMutationVariables>;
export const CreateModulePropagationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModulePropagation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagation"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewModulePropagation"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModulePropagation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagation"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagation"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateModulePropagationMutation, CreateModulePropagationMutationVariables>;
export const ModulePropagationOptionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"modulePropagationOptions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"versions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"variables"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"default"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}}]}}]}}]}}]} as unknown as DocumentNode<ModulePropagationOptionsQuery, ModulePropagationOptionsQueryVariables>;
export const CreateModuleVersionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModuleVersion"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleVersion"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewModuleVersion"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModuleVersion"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"moduleVersion"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleVersion"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateModuleVersionMutation, CreateModuleVersionMutationVariables>;
export const CreateOrgAccountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createOrgAccount"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccount"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewOrgAccount"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createOrgAccount"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgAccount"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccount"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateOrgAccountMutation, CreateOrgAccountMutationVariables>;
export const CreateOrgDimensionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createOrgDimension"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgDimension"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewOrgDimension"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createOrgDimension"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgDimension"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgDimension"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateOrgDimensionMutation, CreateOrgDimensionMutationVariables>;
export const CreateOrgUnitDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createOrgUnit"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgUnit"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewOrgUnit"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createOrgUnit"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgUnit"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgUnit"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateOrgUnitMutation, CreateOrgUnitMutationVariables>;
export const AddAccountToOrgUnitDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"addAccountToOrgUnit"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"addAccountToOrgUnit"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgUnitID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitID"}}},{"kind":"Argument","name":{"kind":"Name","value":"orgAccountID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountID"}}}]}]}}]} as unknown as DocumentNode<AddAccountToOrgUnitMutation, AddAccountToOrgUnitMutationVariables>;
export const OrgDimensionsAndUnitsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"orgDimensionsAndUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]} as unknown as DocumentNode<OrgDimensionsAndUnitsQuery, OrgDimensionsAndUnitsQueryVariables>;
export const OrgAccountsAndMembershipsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"orgAccountsAndMemberships"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimensionID"}}]}}]}}]}}]} as unknown as DocumentNode<OrgAccountsAndMembershipsQuery, OrgAccountsAndMembershipsQueryVariables>;
export const OrgAccountDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"orgAccount"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgAccountID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgAccountID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrgAccountQuery, OrgAccountQueryVariables>;
export const OrgAccountsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"orgAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}}]}}]} as unknown as DocumentNode<OrgAccountsQuery, OrgAccountsQueryVariables>;
export const OrgDimensionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"orgDimension"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgDimensionID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgDimensionID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"parentOrgUnitID"}},{"kind":"Field","name":{"kind":"Name","value":"hierarchy"}}]}}]}}]}}]} as unknown as DocumentNode<OrgDimensionQuery, OrgDimensionQueryVariables>;
export const OrgDimensionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"orgDimensions"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgDimensions"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"limit"},"value":{"kind":"IntValue","value":"100"}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<OrgDimensionsQuery, OrgDimensionsQueryVariables>;
export const OrgUnitDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"orgUnit"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"orgUnit"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"orgUnitID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orgUnitID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"orgDimension"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"upstreamOrgUnits"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccounts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}},{"kind":"Field","name":{"kind":"Name","value":"cloudIdentifier"}}]}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]} as unknown as DocumentNode<OrgUnitQuery, OrgUnitQueryVariables>;
export const PlanExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"planExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"planExecutionRequestID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"planExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"planExecutionRequestID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"planExecutionRequestID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"startedAt"}},{"kind":"Field","name":{"kind":"Name","value":"terraformConfiguration"}},{"kind":"Field","name":{"kind":"Name","value":"moduleAssignment"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"modulePropagation"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"orgAccount"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"cloudPlatform"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleGroup"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"moduleVersion"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"initOutput"}},{"kind":"Field","name":{"kind":"Name","value":"planOutput"}}]}}]}}]} as unknown as DocumentNode<PlanExecutionRequestQuery, PlanExecutionRequestQueryVariables>;
export const CreateModulePropagationDriftCheckRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModulePropagationDriftCheckRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModulePropagationDriftCheckRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationDriftCheckRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"modulePropagationID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationID"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]} as unknown as DocumentNode<CreateModulePropagationDriftCheckRequestMutation, CreateModulePropagationDriftCheckRequestMutationVariables>;
export const CreateModulePropagationExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createModulePropagationExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createModulePropagationExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"modulePropagationExecutionRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"modulePropagationID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"modulePropagationID"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"status"}}]}}]}}]} as unknown as DocumentNode<CreateModulePropagationExecutionRequestMutation, CreateModulePropagationExecutionRequestMutationVariables>;
export const CreateTerraformDriftCheckRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createTerraformDriftCheckRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createTerraformDriftCheckRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"terraformDriftCheckRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"moduleAssignmentID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentID"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateTerraformDriftCheckRequestMutation, CreateTerraformDriftCheckRequestMutationVariables>;
export const CreateTerraformExecutionRequestDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createTerraformExecutionRequest"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"destroy"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Boolean"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createTerraformExecutionRequest"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"terraformExecutionRequest"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"moduleAssignmentID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"moduleAssignmentID"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"destroy"},"value":{"kind":"Variable","name":{"kind":"Name","value":"destroy"}}}]}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateTerraformExecutionRequestMutation, CreateTerraformExecutionRequestMutationVariables>;