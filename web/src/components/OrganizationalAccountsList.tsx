import React, { useState } from "react";
import { OrganizationalAccounts } from "../__generated__/graphql";
import { NavLink, Outlet } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import Container from "react-bootstrap/Container";
import { OrganizationalAccountPage } from "./OrganizationalAccountPage";
import Table from "react-bootstrap/Table";
import { Card, Col, ListGroup, Nav, Row } from "react-bootstrap";
import { NewOrganizationalAccountButton } from "./NewOrganizationalAccountButton";
import { renderCloudPlatform } from "../utils/rendering";

const ORGANIZATIONAL_ACCOUNTS_QUERY = gql`
  query organizationalAccounts {
    organizationalAccounts(limit: 100) {
      items {
        orgAccountId
        name
        cloudPlatform
        cloudIdentifier
      }
    }
  }
`;

type Response = {
  organizationalAccounts: OrganizationalAccounts;
};

export const OrganizationalAccountsList = () => {
  const { loading, error, data, refetch } = useQuery<Response>(
    ORGANIZATIONAL_ACCOUNTS_QUERY,
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
            {data?.organizationalAccounts.items.map((orgAccount) => {
              return (
                <NavLink
                  to={`/org-accounts/${orgAccount?.orgAccountId}`}
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
            <NewOrganizationalAccountButton onCompleted={refetch} />
          </Nav>
        </Col>
        <Col md={"auto"}>
          <Outlet />
        </Col>
      </Row>
    </Container>
  );
};
