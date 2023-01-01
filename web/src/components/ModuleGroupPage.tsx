import React, { useState } from "react";
import { ModuleGroup, ModuleGroups } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Container } from "react-bootstrap";

const MODULE_GROUP_QUERY = gql`
  query moduleGroup($moduleGroupId: ID!) {
    moduleGroup(moduleGroupId: $moduleGroupId) {
      moduleGroupId
      name
      versions {
        items {
          moduleVersionId
          name
          remoteSource
          terraformVersion
        }
      }
      modulePropagations {
        items {
          name
          modulePropagationId
          moduleVersion {
            moduleVersionId
            name
          }
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
  moduleGroup: ModuleGroup;
};

export const ModuleGroupPage = () => {
  const params = useParams();

  const moduleGroupId = params.moduleGroupId ? params.moduleGroupId : "";

  console.log("test");

  const { loading, error, data } = useQuery<Response>(MODULE_GROUP_QUERY, {
    variables: {
      moduleGroupId: moduleGroupId,
    },
  });

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container>
      <h1>
        <b>
          <u>{data?.moduleGroup.name}</u>
        </b>{" "}
        ({data?.moduleGroup.moduleGroupId})<h2></h2>
      </h1>
      <h2>Versions</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Module Version</th>
            <th>Terraform Version</th>
            <th>Remote Source</th>
          </tr>
        </thead>
        <tbody>
          {data?.moduleGroup.versions.items.map((version) => {
            return (
              <tr>
                <td>
                  <NavLink
                    to={`/module-groups/${data?.moduleGroup.moduleGroupId}/versions/${version?.moduleVersionId}`}
                  >
                    {version?.name} ({version?.moduleVersionId})
                  </NavLink>
                </td>
                <td>{version?.terraformVersion}</td>
                <td>{version?.remoteSource}</td>
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
            <th>Module Version</th>
            <th>Org Dimension</th>
            <th>Org Unit</th>
          </tr>
        </thead>
        <tbody>
          {data?.moduleGroup.modulePropagations.items.map((propagation) => {
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
                    to={`/module-groups/${data?.moduleGroup.moduleGroupId}/versions/${propagation?.moduleVersion?.moduleVersionId}`}
                  >
                    {propagation?.moduleVersion?.name} (
                    {propagation?.moduleVersion?.moduleVersionId})
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
