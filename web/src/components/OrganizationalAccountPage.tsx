import React, { useState } from "react";
import {
  OrganizationalAccount,
  OrganizationalAccounts,
} from "../__generated__/graphql";
import { gql } from "../__generated__";
import { NavLink, useParams } from "react-router-dom";
import { useQuery } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Container } from "react-bootstrap";

const ORGANIZATIONAL_ACCOUNT_QUERY = gql(`
  query organizationalAccount($orgAccountId: ID!) {
    organizationalAccount(orgAccountId: $orgAccountId) {
      orgAccountId
      name
      orgUnitMemberships {
        items {
          orgUnit {
            orgUnitId
            name
          }
          orgDimension {
            orgDimensionId
            name
          }
        }
      }
      moduleAccountAssociations {
        items {
          status
          modulePropagationId
          orgAccountId
          modulePropagation {
            moduleGroup {
              moduleGroupId
              name
            }
            moduleVersion {
              moduleVersionId
              name
            }
          }
        }
      }
    }
  }
`);

type Response = {
  organizationalAccount: OrganizationalAccount;
};

export const OrganizationalAccountPage = () => {
  const params = useParams();

  const organizationalAccountId = params.organizationalAccountId
    ? params.organizationalAccountId
    : "";

  console.log("test");

  const { loading, error, data } = useQuery<Response>(
    ORGANIZATIONAL_ACCOUNT_QUERY,
    {
      variables: {
        orgAccountId: organizationalAccountId,
      },
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container>
      <h1>
        <b>
          <u>{data?.organizationalAccount.name}</u>
        </b>{" "}
        ({data?.organizationalAccount.orgAccountId})<h2></h2>
      </h1>
      <h2>Org Unit Memberships</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Org Dimension</th>
            <th>Org Unit</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalAccount.orgUnitMemberships.items.map(
            (orgUnitMembership) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/org-dimensions/${orgUnitMembership?.orgDimension.orgDimensionId}`}
                    >
                      {orgUnitMembership?.orgDimension.name} (
                      {orgUnitMembership?.orgDimension.orgDimensionId})
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/org-dimensions/${orgUnitMembership?.orgDimension.orgDimensionId}/org-units/${orgUnitMembership?.orgUnit.orgUnitId}`}
                    >
                      {orgUnitMembership?.orgUnit.name} (
                      {orgUnitMembership?.orgUnit.orgUnitId})
                    </NavLink>
                  </td>
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
      <h2>Module Propagations</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Module Group</th>
            <th>Module Version</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalAccount.moduleAccountAssociations.items.map(
            (moduleAccountAssociation) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/module-groups/${moduleAccountAssociation?.modulePropagation.moduleGroup.moduleGroupId}`}
                    >
                      {
                        moduleAccountAssociation?.modulePropagation.moduleGroup
                          .name
                      }{" "}
                      (
                      {
                        moduleAccountAssociation?.modulePropagation.moduleGroup
                          .moduleGroupId
                      }
                      )
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/module-groups/${moduleAccountAssociation?.modulePropagation.moduleGroup.moduleGroupId}/module-versions/${moduleAccountAssociation?.modulePropagation.moduleVersion.moduleVersionId}`}
                    >
                      {
                        moduleAccountAssociation?.modulePropagation
                          .moduleVersion.name
                      }{" "}
                      (
                      {
                        moduleAccountAssociation?.modulePropagation
                          .moduleVersion.moduleVersionId
                      }
                      )
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/module-propagations/${moduleAccountAssociation?.modulePropagationId}/account-associations/${moduleAccountAssociation?.orgAccountId}`}
                    >
                      {moduleAccountAssociation?.status}
                    </NavLink>
                  </td>
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
    </Container>
  );
};
