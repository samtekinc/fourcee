import React, { useState } from "react";
import { ModuleVersion, ModuleVersions } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Container } from "react-bootstrap";

const MODULE_VERSION_QUERY = gql`
  query moduleVersion($moduleGroupId: ID!, $moduleVersionId: ID!) {
    moduleVersion(
      moduleGroupId: $moduleGroupId
      moduleVersionId: $moduleVersionId
    ) {
      moduleVersionId
      name
      moduleGroup {
        moduleGroupId
        name
      }
      remoteSource
      terraformVersion
      variables {
        name
        type
        description
        default
      }
      modulePropagations {
        items {
          name
          modulePropagationId
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
`;

type Response = {
  moduleVersion: ModuleVersion;
};

export const ModuleVersionPage = () => {
  const params = useParams();

  const moduleGroupId = params.moduleGroupId ? params.moduleGroupId : "";
  const moduleVersionId = params.moduleVersionId ? params.moduleVersionId : "";

  const { loading, error, data } = useQuery<Response>(MODULE_VERSION_QUERY, {
    variables: {
      moduleGroupId: moduleGroupId,
      moduleVersionId: moduleVersionId,
    },
  });

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container>
      <h1>
        <b>
          <u>{data?.moduleVersion.name}</u>
        </b>{" "}
        ({data?.moduleVersion.moduleVersionId})<h2></h2>
      </h1>
      <p>
        <li>
          <b>Remote Source:</b> {data?.moduleVersion.remoteSource}
        </li>
        <li>
          <b>Terraform Version:</b> {data?.moduleVersion.terraformVersion}
        </li>
      </p>
      <h2>Module Propagations</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Name</th>
            <th>Org Dimension</th>
            <th>Org Unit</th>
          </tr>
        </thead>
        <tbody>
          {data?.moduleVersion.modulePropagations.items.map((propagation) => {
            return (
              <tr>
                <td>
                  <NavLink
                    to={`/module-propagations/${propagation?.modulePropagationId}`}
                  >
                    {propagation?.name} ({propagation?.modulePropagationId})
                  </NavLink>
                </td>
                <td>
                  <NavLink
                    to={`/org-dimensions/${propagation?.orgDimension?.orgDimensionId}`}
                  >
                    {propagation?.orgDimension?.name} (
                    {propagation?.orgDimension?.orgDimensionId})
                  </NavLink>
                </td>
                <td>
                  <NavLink
                    to={`/org-dimensions/${propagation?.orgDimension?.orgDimensionId}/org-units/${propagation?.orgUnit?.orgUnitId}`}
                  >
                    {propagation?.orgUnit?.name} (
                    {propagation?.orgUnit?.orgUnitId})
                  </NavLink>
                </td>
              </tr>
            );
          })}
        </tbody>
      </Table>
    </Container>
  );
};
