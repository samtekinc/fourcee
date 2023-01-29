import { OrgUnit } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Container from "react-bootstrap/Container";
import { Card, Col, ListGroup, Row } from "react-bootstrap";
import { renderCloudPlatform } from "../utils/rendering";
import { NewOrgUnitMembershipButton } from "./NewOrgUnitMembershipButton";
import { DeleteOrgUnitMembershipButton } from "./DeleteOrgUnitMembershipButton";

const ORG_UNIT_QUERY = gql`
  query orgUnit($orgUnitID: ID!) {
    orgUnit(orgUnitID: $orgUnitID) {
      id
      name

      orgDimension {
        id
        name
      }

      upstreamOrgUnits {
        id
        name
        modulePropagations {
          id
          name
          description
          moduleGroup {
            id
            name
          }
          moduleVersion {
            id
            name
          }
        }
      }

      orgAccounts {
        id
        name
        cloudPlatform
        cloudIdentifier
      }

      modulePropagations {
        id
        name
        description
        moduleGroup {
          id
          name
        }
        moduleVersion {
          id
          name
        }
      }
    }
  }
`;

type Response = {
  orgUnit: OrgUnit;
};

export const OrgUnitPage = () => {
  const params = useParams();

  const orgUnitID = params.orgUnitID ? params.orgUnitID : "";
  const orgDimensionID = params.orgDimensionID ? params.orgDimensionID : "";

  const { loading, error, data, refetch } = useQuery<Response>(ORG_UNIT_QUERY, {
    variables: {
      orgUnitID: orgUnitID,
      orgDimensionID: orgDimensionID,
    },
    pollInterval: 10000,
  });

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container style={{ paddingTop: "2rem", borderTop: "1px solid black" }}>
      <h1>
        <b>
          <u>{data?.orgUnit.name}</u>
        </b>{" "}
        ({data?.orgUnit.id})
      </h1>

      <h2>Module Propagations</h2>
      <Row>
        {data?.orgUnit.modulePropagations.map((modulePropagation) => {
          return (
            <Col md={4}>
              <NavLink
                to={`/module-propagations/${modulePropagation?.id}`}
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
        })}
      </Row>

      <br />
      <h2>Upstream Module Propagations</h2>
      <Row>
        {data?.orgUnit.upstreamOrgUnits.map((upstreamOrgUnit) => {
          return (
            <>
              {upstreamOrgUnit?.modulePropagations.map((modulePropagation) => {
                return (
                  <Col md={4}>
                    <NavLink
                      to={`/module-propagations/${modulePropagation?.id}`}
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
                            <b>Module:</b> {modulePropagation?.moduleGroup.name}{" "}
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
              })}
            </>
          );
        })}
      </Row>

      <br />
      <h2>Org Unit Memberships</h2>
      <ListGroup style={{ maxWidth: "36rem" }}>
        {data?.orgUnit.orgAccounts.map((orgAccount) => {
          return (
            <ListGroup.Item>
              <Row>
                <Col md={10}>
                  {renderCloudPlatform(orgAccount.cloudPlatform)}{" "}
                  <NavLink to={`/org-accounts/${orgAccount.id}`}>
                    {orgAccount.name} ({orgAccount.cloudIdentifier})
                  </NavLink>
                </Col>
                <Col md={2}>
                  <DeleteOrgUnitMembershipButton
                    orgUnitID={data?.orgUnit.id}
                    orgAccountID={orgAccount.id}
                    onCompleted={refetch}
                  />
                </Col>
              </Row>
            </ListGroup.Item>
          );
        })}
      </ListGroup>
      <br />
      <NewOrgUnitMembershipButton
        orgDimension={data?.orgUnit.orgDimension}
        orgUnit={data?.orgUnit}
        orgAccount={undefined}
        key={data?.orgUnit.id}
        onCompleted={refetch}
      />
    </Container>
  );
};
