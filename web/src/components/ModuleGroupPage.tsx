import React, { useState } from "react";
import { ModuleGroup, ModuleGroups } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Breadcrumb, Col, Container, Row } from "react-bootstrap";
import {
  renderCloudPlatform,
  renderModuleAssignmentStatus,
} from "../utils/rendering";
import { NewModuleVersionButton } from "./NewModuleVersionButton";

const MODULE_GROUP_QUERY = gql`
  query moduleGroup($moduleGroupId: ID!) {
    moduleGroup(moduleGroupId: $moduleGroupId) {
      moduleGroupId
      cloudPlatform
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
      moduleAssignments {
        items {
          moduleAssignmentId
          moduleVersion {
            moduleVersionId
            name
          }
          modulePropagation {
            modulePropagationId
            name
          }
          orgAccount {
            orgAccountId
            name
          }
          status
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

  const { loading, error, data, refetch } = useQuery<Response>(
    MODULE_GROUP_QUERY,
    {
      variables: {
        moduleGroupId: moduleGroupId,
      },
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container
      style={{ paddingTop: "2rem", maxWidth: "calc(100vw - 20rem)" }}
      fluid
    >
      <Breadcrumb>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/" }}>
          Home
        </Breadcrumb.Item>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/module-groups" }}>
          Modules
        </Breadcrumb.Item>
        <Breadcrumb.Item active>
          {data?.moduleGroup.name} ({data?.moduleGroup.moduleGroupId})
        </Breadcrumb.Item>
      </Breadcrumb>
      <Row>
        <Col md={"auto"}>
          <h1>
            {renderCloudPlatform(data?.moduleGroup.cloudPlatform)}{" "}
            {data?.moduleGroup.name}
          </h1>
        </Col>
      </Row>
      <h2>
        Versions{" "}
        <NewModuleVersionButton
          moduleGroupId={moduleGroupId}
          onCompleted={refetch}
        />
      </h2>
      <Table hover>
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
                    {version?.name}
                  </NavLink>
                </td>
                <td>{version?.terraformVersion}</td>
                <td>{version?.remoteSource}</td>
              </tr>
            );
          })}
        </tbody>
      </Table>
      <Row>
        <Col md={"auto"}>
          <h2>Module Propagations</h2>
          <Table hover>
            <thead>
              <tr>
                <th>Name</th>
                <th>Module Version</th>
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
                        {propagation?.name}
                      </NavLink>
                    </td>
                    <td>
                      <NavLink
                        to={`/module-groups/${data?.moduleGroup.moduleGroupId}/versions/${propagation?.moduleVersion?.moduleVersionId}`}
                      >
                        {propagation?.moduleVersion?.name}
                      </NavLink>
                    </td>
                    <td>
                      <NavLink
                        to={`/org-dimensions/${propagation?.orgDimension?.orgDimensionId}`}
                      >
                        {propagation?.orgDimension?.name}
                      </NavLink>
                      {" / "}
                      <NavLink
                        to={`/org-dimensions/${propagation?.orgDimension?.orgDimensionId}/org-units/${propagation?.orgUnit?.orgUnitId}`}
                      >
                        {propagation?.orgUnit?.name}
                      </NavLink>
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </Table>
        </Col>
        <Col md={"auto"}>
          <h2>Module Assignments</h2>
          <Table hover>
            <thead>
              <tr>
                <th>Assignment Id</th>
                <th>Account</th>
                <th>Module Version</th>
                <th>Status</th>
                <th>Propagated By</th>
              </tr>
            </thead>
            <tbody>
              {data?.moduleGroup.moduleAssignments.items.map((assignment) => {
                return (
                  <tr>
                    <td>
                      <NavLink
                        to={`/module-assignments/${assignment?.moduleAssignmentId}`}
                      >
                        {assignment?.moduleAssignmentId}
                      </NavLink>
                    </td>
                    <td>
                      <NavLink
                        to={`/org-accounts/${assignment?.orgAccount?.orgAccountId}`}
                      >
                        {assignment?.orgAccount?.name}
                      </NavLink>
                    </td>
                    <td>
                      <NavLink
                        to={`/module-groups/${moduleGroupId}/versions/${assignment?.moduleVersion?.moduleVersionId}`}
                      >
                        {assignment?.moduleVersion?.name}
                      </NavLink>
                    </td>
                    <td>{renderModuleAssignmentStatus(assignment?.status)}</td>
                    <td>
                      {assignment?.modulePropagation ? (
                        <NavLink
                          to={`/module-propagations/${assignment?.modulePropagation.modulePropagationId}`}
                        >
                          {assignment?.modulePropagation?.name}
                        </NavLink>
                      ) : (
                        <div>Direct Assignment</div>
                      )}
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </Table>
        </Col>
      </Row>
    </Container>
  );
};
