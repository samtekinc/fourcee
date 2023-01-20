import React, { useState } from "react";
import { ModuleVersion, ModuleVersions } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Breadcrumb, Card, Col, Container, Row } from "react-bootstrap";
import { renderCloudPlatform, renderRemoteSource } from "../utils/rendering";

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
        cloudPlatform
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
          description
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
        <Breadcrumb.Item
          linkAs={NavLink}
          linkProps={{
            to: `/module-groups/${data?.moduleVersion.moduleGroup.moduleGroupId}`,
          }}
        >
          {data?.moduleVersion.moduleGroup.name} (
          {data?.moduleVersion.moduleGroup.moduleGroupId})
        </Breadcrumb.Item>
        <Breadcrumb.Item
          linkAs={NavLink}
          linkProps={{
            to: `/module-groups/${data?.moduleVersion.moduleGroup.moduleGroupId}`,
          }}
        >
          Versions
        </Breadcrumb.Item>
        <Breadcrumb.Item active>
          {data?.moduleVersion.name} ({data?.moduleVersion.moduleVersionId})
        </Breadcrumb.Item>
      </Breadcrumb>

      <Row>
        <Col md={"auto"}>
          <h1>
            {renderCloudPlatform(data?.moduleVersion.moduleGroup.cloudPlatform)}{" "}
            {data?.moduleVersion.moduleGroup.name} {data?.moduleVersion.name}
          </h1>
        </Col>
      </Row>

      <p>
        <b>Remote Source:</b>{" "}
        {renderRemoteSource(data?.moduleVersion.remoteSource)}
        <br />
        <b>Terraform Version:</b> {data?.moduleVersion.terraformVersion}
      </p>
      <h2>Variables</h2>
      <Container fluid style={{ maxHeight: "50vh", overflow: "auto" }}>
        <Table striped hover bordered responsive>
          <thead>
            <tr>
              <th>Name</th>
              <th>Type</th>
              <th>Default</th>
              <th>Description</th>
            </tr>
          </thead>
          <tbody>
            {data?.moduleVersion.variables.map((variable) => {
              return (
                <tr>
                  <td>{variable?.name}</td>
                  <td>{variable?.type}</td>
                  <td>{variable?.default}</td>
                  <td>{variable?.description}</td>
                </tr>
              );
            })}
          </tbody>
        </Table>
      </Container>
      <br />
      <h2>Module Propagations</h2>
      <Row>
        {data?.moduleVersion.modulePropagations.items.map((propagation) => {
          return (
            <Col md={"auto"}>
              <Card>
                <Card.Body>
                  <NavLink
                    to={`/module-propagations/${propagation?.modulePropagationId}`}
                  >
                    <Card.Title style={{ fontSize: "medium" }}>
                      {propagation?.name}
                    </Card.Title>
                  </NavLink>
                  <Card.Text style={{ fontSize: "small" }}>
                    <b>Org Unit: </b>
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
                    <br />
                    <b>Description: </b>
                    {propagation?.description}
                  </Card.Text>
                </Card.Body>
              </Card>
            </Col>
          );
        })}
      </Row>
    </Container>
  );
};
