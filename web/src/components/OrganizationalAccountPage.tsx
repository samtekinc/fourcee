import React, { useState } from "react";
import {
  OrganizationalAccount,
  OrganizationalAccounts,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Breadcrumb, Card, Col, Container, Row } from "react-bootstrap";
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
    <Container
      style={{ paddingTop: "2rem", maxWidth: "calc(100vw - 20rem)" }}
      fluid
    >
      <Breadcrumb>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/" }}>
          Home
        </Breadcrumb.Item>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/module-groups" }}>
          Accounts
        </Breadcrumb.Item>
        <Breadcrumb.Item active>
          {data?.organizationalAccount.name} (
          {data?.organizationalAccount.orgAccountId})
        </Breadcrumb.Item>
      </Breadcrumb>
      <Row>
        <Col md={"auto"}>
          <h1>
            {renderCloudPlatform(data?.organizationalAccount.cloudPlatform)}{" "}
            {data?.organizationalAccount.name}
          </h1>
        </Col>
      </Row>
      <p>
        <b>Cloud Identifier:</b> {data?.organizationalAccount.cloudIdentifier}
      </p>

      <Row>
        <Col md={"auto"}>
          <h2>Org Unit Memberships</h2>
          <Table hover>
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
        </Col>
      </Row>
      <br />

      <h2>Module Assignments</h2>
      <Row>
        {data?.organizationalAccount.moduleAssignments.items.map(
          (moduleAssignment) => {
            return (
              <Col md={"auto"} style={{ padding: "0.25rem" }}>
                <Card>
                  <Card.Body>
                    <NavLink
                      to={`/module-assignments/${moduleAssignment?.moduleAssignmentId}`}
                    >
                      <Card.Title>
                        {moduleAssignment?.moduleAssignmentId}{" "}
                      </Card.Title>
                    </NavLink>
                    <Card.Text>
                      {renderModuleAssignmentStatus(moduleAssignment?.status)}
                      <br />
                      <b>Module: </b>
                      <NavLink
                        to={`/module-groups/${moduleAssignment?.moduleGroup.moduleGroupId}`}
                      >
                        {moduleAssignment?.moduleGroup.name}
                      </NavLink>
                      {" / "}
                      <NavLink
                        to={`/module-groups/${moduleAssignment?.moduleGroup.moduleGroupId}/module-versions/${moduleAssignment?.moduleVersion.moduleVersionId}`}
                      >
                        {moduleAssignment?.moduleVersion.name}
                      </NavLink>
                      <br />
                      <b>Propagated By: </b>
                      {moduleAssignment?.modulePropagation ? (
                        <>
                          <NavLink
                            to={`/module-propagations/${moduleAssignment?.modulePropagation?.modulePropagationId}`}
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
                        <span>Direct Assignment</span>
                      )}
                    </Card.Text>
                  </Card.Body>
                </Card>
              </Col>
            );
          }
        )}
      </Row>
    </Container>
  );
};
