import { OrgAccount } from "../__generated__/graphql";
import { NavLink, Outlet } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Container from "react-bootstrap/Container";
import { Card, Col, ListGroup, Nav, Row } from "react-bootstrap";
import { NewOrgAccountButton } from "./NewOrgAccountButton";
import { renderCloudPlatform } from "../utils/rendering";

const ORG_ACCOUNTS_QUERY = gql`
  query orgAccounts {
    orgAccounts {
      id
      name
      cloudPlatform
      cloudIdentifier
    }
  }
`;

type Response = {
  orgAccounts: OrgAccount[];
};

export const OrgAccountsList = () => {
  const { loading, error, data, refetch } = useQuery<Response>(
    ORG_ACCOUNTS_QUERY,
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
          <h3>Org Accounts</h3>
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
            {data?.orgAccounts.map((orgAccount) => {
              return (
                <NavLink
                  to={`/org-accounts/${orgAccount?.id}`}
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
                      <Card.Title style={{ fontSize: "medium" }}>
                        {renderCloudPlatform(orgAccount?.cloudPlatform)}{" "}
                        {orgAccount?.name}
                      </Card.Title>
                      <Card.Text style={{ fontSize: "small" }}>
                        <b>Cloud ID: </b>
                        {orgAccount?.cloudIdentifier}
                      </Card.Text>
                    </Card.Body>
                  </Card>
                </NavLink>
              );
            })}
            <br />
            <NewOrgAccountButton onCompleted={refetch} />
          </Nav>
        </Col>
        <Col md={"auto"}>
          <Outlet />
        </Col>
      </Row>
    </Container>
  );
};
