import React, { useState } from "react";
import {
  OrganizationalUnit,
  OrganizationalUnits,
} from "../__generated__/graphql";
import { gql } from "../__generated__";
import { NavLink, useParams } from "react-router-dom";
import { useQuery } from "@apollo/client";
import Table from "react-bootstrap/Table";
import Container from "react-bootstrap/Container";

const ORGANIZATIONAL_UNIT_QUERY = gql(`
  query organizationalUnit($orgUnitId: ID!, $orgDimensionId: ID!) {
    organizationalUnit(orgDimensionId: $orgDimensionId, orgUnitId: $orgUnitId) {
    orgUnitId
    orgDimensionId
    name
    hierarchy
    parentOrgUnitId
    children {
      items {
        orgUnitId
        name
        hierarchy
      }
    }
    downstreamOrgUnits {
      items {
        orgUnitId
        name
        hierarchy
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
  }
}
`);

type Response = {
  organizationalUnit: OrganizationalUnit;
};

export const OrganizationalUnitPage = () => {
  const params = useParams();

  const organizationalUnitId = params.organizationalUnitId
    ? params.organizationalUnitId
    : "";
  const organizationalDimensionId = params.organizationalDimensionId
    ? params.organizationalDimensionId
    : "";

  const { loading, error, data } = useQuery<Response>(
    ORGANIZATIONAL_UNIT_QUERY,
    {
      variables: {
        orgUnitId: organizationalUnitId,
        orgDimensionId: organizationalDimensionId,
      },
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container>
      <h1>
        <b>
          <u>{data?.organizationalUnit.name}</u>
        </b>{" "}
        ({data?.organizationalUnit.orgUnitId})
      </h1>
      <h2>Children</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Org Dimension Id</th>
            <th>Name</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalUnit.children.items.map((orgUnit) => {
            return (
              <tr>
                <td>
                  <NavLink
                    to={`/org-dimensions/${organizationalDimensionId}/org-units/${orgUnit?.orgUnitId}`}
                  >
                    {orgUnit?.orgUnitId}
                  </NavLink>
                </td>
                <td>{orgUnit?.name}</td>
              </tr>
            );
          })}
        </tbody>
      </Table>
      <h2>Downstream Org Units</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Org Dimension Id</th>
            <th>Name</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalUnit.downstreamOrgUnits.items.map((orgUnit) => {
            return (
              <tr>
                <td>
                  <NavLink
                    to={`/org-dimensions/${organizationalDimensionId}/org-units/${orgUnit?.orgUnitId}`}
                  >
                    {orgUnit?.orgUnitId}
                  </NavLink>
                </td>
                <td>{orgUnit?.name}</td>
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
            <th>Module Version Id</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalUnit.modulePropagations.items.map(
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
            <th>Org Account Id</th>
            <th>Org Account Name</th>
            <th>Cloud Platform</th>
            <th>Cloud ID</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalUnit.orgUnitMemberships.items.map(
            (membership) => {
              return (
                <tr>
                  <td>
                    <NavLink
                      to={`/org-accounts/${membership?.orgAccount.orgAccountId}`}
                    >
                      {membership?.orgAccount.orgAccountId}
                    </NavLink>
                  </td>
                  <td>{membership?.orgAccount.name}</td>
                  <td>{membership?.orgAccount.cloudPlatform}</td>
                  <td>{membership?.orgAccount.cloudIdentifier}</td>
                </tr>
              );
            }
          )}
        </tbody>
      </Table>
    </Container>
  );
};
