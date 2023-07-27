import { OrgAccount } from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import { Breadcrumb, Card, Col, Container, Row } from "react-bootstrap";
import {
  renderCloudPlatform,
  renderModuleAssignmentStatus,
} from "../utils/rendering";

const ORG_ACCOUNT_QUERY = gql`
  query orgAccount($orgAccountID: ID!) {
    orgAccount(orgAccountID: $orgAccountID) {
      id
      name
      cloudPlatform
      cloudIdentifier
      orgUnits {
        id
        name
        orgDimension {
          id
          name
        }
      }
      moduleAssignments {
        id
        name
        status
        moduleGroup {
          id
          name
        }
        moduleVersion {
          id
          name
        }
        modulePropagation {
          id
          name
          orgUnit {
            id
            name
          }
          orgDimension {
            id
            name
          }
        }
      }
    }
  }
`;

type Response = {
  orgAccount: OrgAccount;
};

export const OrgAccountPage = () => {
  const params = useParams();

  const orgAccountID = params.orgAccountID ? params.orgAccountID : "";

  const { loading, error, data } = useQuery<Response>(ORG_ACCOUNT_QUERY, {
    variables: {
      orgAccountID: orgAccountID,
    },
    pollInterval: 5000,
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
          Accounts
        </Breadcrumb.Item>
        <Breadcrumb.Item active>
          {data?.orgAccount.name} ({data?.orgAccount.id})
        </Breadcrumb.Item>
      </Breadcrumb>
      <Row>
        <Col md={"auto"}>
          <h1>
            {renderCloudPlatform(data?.orgAccount.cloudPlatform)}{" "}
            {data?.orgAccount.name}
          </h1>
        </Col>
      </Row>
      <p>
        <b>Cloud Identifier:</b> {data?.orgAccount.cloudIdentifier}
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
              {data?.orgAccount.orgUnits.map((orgUnit) => {
                return (
                  <tr>
                    <td>
                      <NavLink
                        to={`/org-structures/${orgUnit?.orgDimension.id}`}
                      >
                        {orgUnit?.orgDimension.name}
                      </NavLink>
                    </td>
                    <td>
                      <NavLink
                        to={`/org-structures/${orgUnit?.orgDimension.id}/org-units/${orgUnit?.id}`}
                      >
                        {orgUnit?.name}
                      </NavLink>
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </Table>
        </Col>
      </Row>
      <br />

      <h2>Module Assignments</h2>
      <Row>
        {data?.orgAccount.moduleAssignments.map((moduleAssignment) => {
          return (
            <Col md={"auto"} style={{ padding: "0.25rem" }}>
              <Card>
                <Card.Body>
                  <NavLink to={`/module-assignments/${moduleAssignment?.id}`}>
                    <Card.Title>{moduleAssignment?.name} </Card.Title>
                  </NavLink>
                  <Card.Text>
                    {renderModuleAssignmentStatus(moduleAssignment?.status)}
                    <br />
                    <b>Module: </b>
                    <NavLink
                      to={`/module-groups/${moduleAssignment?.moduleGroup.id}`}
                    >
                      {moduleAssignment?.moduleGroup.name}
                    </NavLink>
                    {" / "}
                    <NavLink
                      to={`/module-groups/${moduleAssignment?.moduleGroup.id}/module-versions/${moduleAssignment?.moduleVersion.id}`}
                    >
                      {moduleAssignment?.moduleVersion.name}
                    </NavLink>
                    <br />
                    <b>Propagated By: </b>
                    {moduleAssignment?.modulePropagation ? (
                      <>
                        <NavLink
                          to={`/module-propagations/${moduleAssignment?.modulePropagation?.id}`}
                        >
                          {moduleAssignment?.modulePropagation?.name}
                        </NavLink>
                        {" ("}
                        <NavLink
                          to={`/org-structures/${moduleAssignment?.modulePropagation?.orgDimension.id}`}
                        >
                          {
                            moduleAssignment?.modulePropagation?.orgDimension
                              .name
                          }
                        </NavLink>
                        {" / "}
                        <NavLink
                          to={`/org-structures/${moduleAssignment?.modulePropagation?.orgDimension.id}/org-units/${moduleAssignment?.modulePropagation?.orgUnit.id}`}
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
        })}
      </Row>
    </Container>
  );
};
