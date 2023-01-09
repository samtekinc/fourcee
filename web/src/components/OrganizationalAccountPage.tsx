import React, { useState } from "react";
import {
  OrganizationalAccount,
  OrganizationalAccounts,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Container } from "react-bootstrap";

const ORGANIZATIONAL_ACCOUNT_QUERY = gql`
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
      moduleAssignments {
        items {
          moduleAssignmentId
          status
          modulePropagationId
          orgAccountId
          moduleGroup {
            moduleGroupId
            name
          }
          moduleVersion {
            moduleVersionId
            name
          }
          modulePropagation {
            name
          }
        }
      }
    }
  }
`;

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
            <th>Module Propagation</th>
            <th>Module Group</th>
            <th>Module Version</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalAccount.moduleAssignments.items.map(
            (moduleAssignment) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/module-propagations/${moduleAssignment?.modulePropagationId}`}
                    >
                      {moduleAssignment?.modulePropagation?.name} (
                      {moduleAssignment?.modulePropagationId})
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/module-groups/${moduleAssignment?.moduleGroup.moduleGroupId}`}
                    >
                      {moduleAssignment?.moduleGroup.name} (
                      {moduleAssignment?.moduleGroup.moduleGroupId})
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/module-groups/${moduleAssignment?.moduleGroup.moduleGroupId}/module-versions/${moduleAssignment?.moduleVersion.moduleVersionId}`}
                    >
                      {moduleAssignment?.moduleVersion.name} (
                      {moduleAssignment?.moduleVersion.moduleVersionId})
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/module-assignments/${moduleAssignment?.moduleAssignmentId}`}
                    >
                      {moduleAssignment?.status}
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
