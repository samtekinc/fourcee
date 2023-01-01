import React, { useState } from "react";
import {
  OrganizationalDimension,
  OrganizationalDimensions,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import Container from "react-bootstrap/Container";

const ORGANIZATIONAL_DIMENSION_QUERY = gql`
  query organizationalDimension($orgDimensionId: ID!) {
    organizationalDimension(orgDimensionId: $orgDimensionId) {
      orgDimensionId
      name
      orgUnits {
        items {
          orgUnitId
          name
          parentOrgUnitId
          hierarchy
        }
      }
      modulePropagations {
        items {
          modulePropagationId
          moduleGroupId
          moduleVersionId
          orgUnitId
          orgDimensionId
          name
          description
        }
      }
      orgUnitMemberships {
        items {
          orgAccount {
            orgAccountId
            name
            cloudPlatform
            cloudIdentifier
          }
          orgUnit {
            orgUnitId
            name
          }
        }
      }
    }
  }
`;

type Response = {
  organizationalDimension: OrganizationalDimension;
};

export const OrganizationalDimensionPage = () => {
  const params = useParams();

  const organizationalDimensionId = params.organizationalDimensionId
    ? params.organizationalDimensionId
    : "";

  const { loading, error, data } = useQuery<Response>(
    ORGANIZATIONAL_DIMENSION_QUERY,
    {
      variables: {
        orgDimensionId: organizationalDimensionId,
      },
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  let orgUnitsSorted = Array.prototype.slice
    .call(data?.organizationalDimension.orgUnits.items)
    .sort((a, b) =>
      a.hierarchy + a.orgUnitId > b.hierarchy + b.orgUnitId ? 1 : -1
    );

  return (
    <Container>
      <h1>
        <b>
          <u>{data?.organizationalDimension.name}</u>
        </b>{" "}
        ({data?.organizationalDimension.orgDimensionId})
      </h1>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Org Unit Name</th>
            <th>ID</th>
            <th>Parent</th>
            <th>Hierarchy</th>
          </tr>
        </thead>
        <tbody>
          {orgUnitsSorted?.map((orgUnit) => {
            return (
              <tr>
                <td>
                  <NavLink
                    to={`/org-dimensions/${organizationalDimensionId}/org-units/${orgUnit.orgUnitId}`}
                  >
                    {orgUnit.name}
                  </NavLink>
                </td>
                <td>{orgUnit.orgUnitId}</td>
                <td>{orgUnit.parentOrgUnitId}</td>
                <td>{orgUnit.hierarchy}</td>
              </tr>
            );
          })}
        </tbody>
      </Table>
      <h2>Module Propagations</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Name</th>
            <th>ID</th>
            <th>Org Unit Id</th>
            <th>Module Version Id</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalDimension.modulePropagations.items.map(
            (modulePropagation) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/module-propagations/${modulePropagation?.modulePropagationId}`}
                    >
                      {modulePropagation?.name}
                    </NavLink>
                  </td>
                  <td>{modulePropagation?.modulePropagationId}</td>
                  <td>{modulePropagation?.orgUnitId}</td>
                  <td>{modulePropagation?.moduleVersionId}</td>
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
      <h2>Org Unit Memberships</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Org Account</th>
            <th>Cloud Platform</th>
            <th>Cloud ID</th>
            <th>Org Unit</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalDimension.orgUnitMemberships.items.map(
            (membership) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/org-accounts/${membership?.orgAccount.orgAccountId}`}
                    >
                      {membership?.orgAccount.name} (
                      {membership?.orgAccount.orgAccountId})
                    </NavLink>
                  </td>
                  <td>{membership?.orgAccount.cloudPlatform}</td>
                  <td>{membership?.orgAccount.cloudIdentifier}</td>
                  <td>
                    <NavLink
                      to={`/org-dimensions/${organizationalDimensionId}/org-units/${membership?.orgUnit.orgUnitId}`}
                    >
                      {membership?.orgUnit.name} (
                      {membership?.orgUnit.orgUnitId})
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
