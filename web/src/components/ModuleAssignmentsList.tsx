import { ModuleAssignment } from "../__generated__/graphql";
import { NavLink, Outlet } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Container from "react-bootstrap/Container";
import { Card, Col, ListGroup, Nav, Row } from "react-bootstrap";
import { NewModuleAssignmentButton } from "./NewModuleAssignmentButton";
import { renderCloudPlatform } from "../utils/rendering";

const MODULE_ASSIGNMENTS_QUERY = gql`
  query moduleAssignments {
    unpropagated: moduleAssignments(filters: { isPropagated: false }) {
      id
      name
      description
      moduleGroup {
        id
        name
        cloudPlatform
      }
      moduleVersion {
        id
        name
      }
      orgAccount {
        id
        name
        cloudPlatform
      }
    }
  }
`;

type Response = {
  unpropagated: ModuleAssignment[];
};

export const ModuleAssignmentsList = () => {
  const { loading, error, data, refetch } = useQuery<Response>(
    MODULE_ASSIGNMENTS_QUERY,
    {
      pollInterval: 30000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container fluid>
      <Row className="flex-xl-nowrap">
        <Col
          md={2}
          xl={2}
          className="d-flex flex-column"
          style={{
            position: "sticky",
            top: "3.5rem",
            backgroundColor: "#f7f7f7",
            zIndex: 1000,
            maxWidth: "20rem",
            height: "calc(100vh - 3.5rem)",
            borderRight: "1px solid #dee2e6",
            paddingTop: "1rem",
          }}
        >
          <h3>Module Assignments</h3>
          <Nav
            as={ListGroup}
            style={{
              maxHeight: "calc(100vh - 7rem)",
              flexDirection: "column",
              height: "100%",
              display: "flex",
              overflow: "auto",
              flexWrap: "nowrap",
            }}
          >
            {data?.unpropagated.map((moduleAssignment) => {
              return (
                <NavLink
                  to={`/module-assignments/${moduleAssignment?.id}`}
                  style={({ isActive }) =>
                    isActive
                      ? {
                          color: "blue",
                          textDecoration: "none",
                          padding: "0.25rem",
                        }
                      : {
                          color: "inherit",
                          textDecoration: "none",
                          padding: "0.25rem",
                        }
                  }
                >
                  <Card>
                    <Card.Body>
                      <Card.Title style={{ fontSize: "small" }}>
                        {renderCloudPlatform(
                          moduleAssignment?.orgAccount.cloudPlatform
                        )}{" "}
                        {moduleAssignment?.name}
                      </Card.Title>
                      <Card.Text style={{ fontSize: "x-small" }}>
                        <b>Module: </b>
                        {moduleAssignment?.moduleGroup?.name}
                        {" / "}
                        {moduleAssignment?.moduleVersion?.name}

                        <br />
                        <b>Org Account: </b>
                        {moduleAssignment?.orgAccount?.name}
                        <br />
                        <b>Description: </b>
                        {moduleAssignment?.description}
                      </Card.Text>
                    </Card.Body>
                  </Card>
                </NavLink>
              );
            })}
            <br />
            <NewModuleAssignmentButton onCompleted={refetch} />
            <br />
          </Nav>
        </Col>
        <Col md={"auto"}>
          <Outlet />
        </Col>
      </Row>
    </Container>
  );
};
