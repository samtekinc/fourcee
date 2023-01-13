import React, { useState } from "react";
import {
  OrganizationalAccount,
  OrganizationalAccounts,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Container } from "react-bootstrap";
import {
  renderCloudPlatform,
  renderModuleAssignmentStatus,
} from "../utils/rendering";

const ORGANIZATIONAL_ACCOUNT_QUERY = gql`
  query organizationalAccount($orgAccountId: ID!) {
    organizationalAccount(orgAccountId: $orgAccountId) {
      orgAccountId
      name
      cloudPlatform
      cloudIdentifier
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
            modulePropagationId
            name
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

  const { loading, error, data } = useQuery<Response>(
    ORGANIZATIONAL_ACCOUNT_QUERY,
    {
      variables: {
        orgAccountId: organizationalAccountId,
      },
      pollInterval: 5000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container>
      <h1>
        {renderCloudPlatform(data?.organizationalAccount.cloudPlatform)}{" "}
        <i>{data?.organizationalAccount.name}</i> (
        {data?.organizationalAccount.orgAccountId})<h2></h2>
      </h1>
      <p>
        <b>Cloud Identifier:</b> {data?.organizationalAccount.cloudIdentifier}
      </p>
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
                      {orgUnitMembership?.orgDimension.name}
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/org-dimensions/${orgUnitMembership?.orgDimension.orgDimensionId}/org-units/${orgUnitMembership?.orgUnit.orgUnitId}`}
                    >
                      {orgUnitMembership?.orgUnit.name}
                    </NavLink>
                  </td>
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
      <h2>Module Assignments</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Status</th>
            <th>Assignment Id</th>
            <th>Module Group</th>
            <th>Module Version</th>
            <th>Propagated By</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalAccount.moduleAssignments.items.map(
            (moduleAssignment) => {
              return (
                <tr>
                  <td>
                    {renderModuleAssignmentStatus(moduleAssignment?.status)}
                  </td>
                  <td>
                    <NavLink
                      to={`/module-assignments/${moduleAssignment?.moduleAssignmentId}`}
                    >
                      {moduleAssignment?.moduleAssignmentId}
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/module-groups/${moduleAssignment?.moduleGroup.moduleGroupId}`}
                    >
                      {moduleAssignment?.moduleGroup.name}
                    </NavLink>
                  </td>
                  <td>
                    <NavLink
                      to={`/module-groups/${moduleAssignment?.moduleGroup.moduleGroupId}/module-versions/${moduleAssignment?.moduleVersion.moduleVersionId}`}
                    >
                      {moduleAssignment?.moduleVersion.name}
                    </NavLink>
                  </td>
                  <td>
                    {moduleAssignment?.modulePropagation ? (
                      <>
                        <NavLink
                          to={`/module-propagations/${moduleAssignment?.modulePropagation.modulePropagationId}`}
                        >
                          {moduleAssignment?.modulePropagation?.name}
                        </NavLink>
                        {" ("}
                        <NavLink
                          to={`/org-dimensions/${moduleAssignment?.modulePropagation?.orgDimension.orgDimensionId}`}
                        >
                          {
                            moduleAssignment?.modulePropagation?.orgDimension
                              .name
                          }
                        </NavLink>
                        {" / "}
                        <NavLink
                          to={`/org-dimensions/${moduleAssignment?.modulePropagation?.orgDimension.orgDimensionId}/org-units/${moduleAssignment?.modulePropagation?.orgUnit.orgUnitId}`}
                        >
                          {moduleAssignment?.modulePropagation?.orgUnit.name}
                        </NavLink>
                        {")"}
                      </>
                    ) : (
                      <div>Direct Assignment</div>
                    )}
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
