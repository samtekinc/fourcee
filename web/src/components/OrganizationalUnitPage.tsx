import React, { useState } from "react";
import {
  OrganizationalUnit,
  OrganizationalUnits,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import Container from "react-bootstrap/Container";
import { GetOrgUnitTree, OrgUnitTreeNode } from "../utils/org_tree_rendering";
import { Tree } from "react-organizational-chart";
import { Card, CardGroup, Col, ListGroup, Row } from "react-bootstrap";
import { renderCloudPlatform } from "../utils/rendering";
import { NewOrganizationalUnitMembershipButton } from "./NewOrganizationalUnitMembershipButton";
import { DeleteOrganizationalUnitMembershipButton } from "./DeleteOrganizationalUnitMembershipButton";

const ORGANIZATIONAL_UNIT_QUERY = gql`
  query organizationalUnit($orgUnitId: ID!, $orgDimensionId: ID!) {
    organizationalUnit(orgDimensionId: $orgDimensionId, orgUnitId: $orgUnitId) {
      orgUnitId
      orgDimension {
        orgDimensionId
        name
      }
      name
      hierarchy
      parentOrgUnitId
      upstreamOrgUnits {
        items {
          orgUnitId
          name
          modulePropagations {
            items {
              modulePropagationId
              name
              description
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
      orgUnitMemberships {
        items {
          orgDimensionId
          orgAccountId
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
          name
          description
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
`;

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

  const { loading, error, data, refetch } = useQuery<Response>(
    ORGANIZATIONAL_UNIT_QUERY,
    {
      variables: {
        orgUnitId: organizationalUnitId,
        orgDimensionId: organizationalDimensionId,
      },
      pollInterval: 10000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container style={{ paddingTop: "2rem", borderTop: "1px solid black" }}>
      <h1>
        <b>
          <u>{data?.organizationalUnit.name}</u>
        </b>{" "}
        ({data?.organizationalUnit.orgUnitId})
      </h1>

      <h2>Module Propagations</h2>
      <Row>
        {data?.organizationalUnit.modulePropagations.items.map(
          (modulePropagation) => {
            return (
              <Col md={4}>
                <NavLink
                  to={`/module-propagations/${modulePropagation?.modulePropagationId}`}
                  style={{ textDecoration: "none", color: "black" }}
                >
                  <Card>
                    <Card.Body>
                      <Card.Title style={{ fontSize: "large", color: "blue" }}>
                        {modulePropagation?.name}
                      </Card.Title>
                      <Card.Text style={{ fontSize: "small" }}>
                        <b>Module:</b> {modulePropagation?.moduleGroup.name}{" "}
                        {modulePropagation?.moduleVersion.name}
                      </Card.Text>
                      <Card.Text style={{ fontSize: "small" }}>
                        {modulePropagation?.description}
                      </Card.Text>
                    </Card.Body>
                  </Card>
                </NavLink>
                <br />
              </Col>
            );
          }
        )}
      </Row>

      <br />
      <h2>Upstream Module Propagations</h2>
      <Row>
        {data?.organizationalUnit.upstreamOrgUnits.items.map(
          (upstreamOrgUnit) => {
            return (
              <>
                {upstreamOrgUnit?.modulePropagations.items.map(
                  (modulePropagation) => {
                    return (
                      <Col md={4}>
                        <NavLink
                          to={`/module-propagations/${modulePropagation?.modulePropagationId}`}
                          style={{ textDecoration: "none", color: "black" }}
                        >
                          <Card>
                            <Card.Body>
                              <Card.Title
                                style={{ fontSize: "large", color: "blue" }}
                              >
                                {modulePropagation?.name}
                              </Card.Title>
                              <Card.Text style={{ fontSize: "small" }}>
                                <b>Module:</b>{" "}
                                {modulePropagation?.moduleGroup.name}{" "}
                                {modulePropagation?.moduleVersion.name}
                                <br />
                                <b>Propagated By:</b> {upstreamOrgUnit?.name}
                              </Card.Text>
                              <Card.Text style={{ fontSize: "small" }}>
                                {modulePropagation?.description}
                              </Card.Text>
                            </Card.Body>
                          </Card>
                        </NavLink>
                        <br />
                      </Col>
                    );
                  }
                )}
              </>
            );
          }
        )}
      </Row>

      <br />
      <h2>Org Unit Memberships</h2>
      <ListGroup style={{ maxWidth: "36rem" }}>
        {data?.organizationalUnit.orgUnitMemberships.items.map((membership) => {
          return (
            <ListGroup.Item>
              <Row>
                <Col md={10}>
                  {renderCloudPlatform(membership?.orgAccount.cloudPlatform)}{" "}
                  <NavLink
                    to={`/org-accounts/${membership?.orgAccount.orgAccountId}`}
                  >
                    {membership?.orgAccount.name} (
                    {membership?.orgAccount.cloudIdentifier})
                  </NavLink>
                </Col>
                <Col md={2}>
                  <DeleteOrganizationalUnitMembershipButton
                    orgDimensionId={membership?.orgDimensionId}
                    orgAccountId={membership?.orgAccountId}
                    onCompleted={refetch}
                  />
                </Col>
              </Row>
            </ListGroup.Item>
          );
        })}
      </ListGroup>
      <br />
      <NewOrganizationalUnitMembershipButton
        orgDimension={data?.organizationalUnit.orgDimension}
        orgUnit={data?.organizationalUnit}
        orgAccount={undefined}
        key={data?.organizationalUnit.orgUnitId}
        onCompleted={refetch}
      />
    </Container>
  );
};
