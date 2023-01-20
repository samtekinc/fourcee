import React, { useState } from "react";
import { OrganizationalDimensions } from "../__generated__/graphql";
import { NavLink, Outlet } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import Container from "react-bootstrap/Container";
import { OrganizationalDimensionPage } from "./OrganizationalDimensionPage";
import Table from "react-bootstrap/Table";
import { Card, Col, ListGroup, Nav, Row } from "react-bootstrap";
import { NewOrganizationalDimensionButton } from "./NewOrganizationalDimensionButton";

const ORGANIZATIONAL_DIMENSIONS_QUERY = gql`
  query organizationalDimensions {
    organizationalDimensions(limit: 100) {
      items {
        orgDimensionId
        name
        orgUnits {
          items {
            orgUnitId
          }
        }
        orgUnitMemberships {
          items {
            orgAccountId
          }
        }
      }
    }
  }
`;

type Response = {
  organizationalDimensions: OrganizationalDimensions;
};

export const OrganizationalDimensionsList = () => {
  const { loading, error, data, refetch } = useQuery<Response>(
    ORGANIZATIONAL_DIMENSIONS_QUERY,
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
          <h3>Org Dimensions</h3>
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
            {data?.organizationalDimensions.items.map((orgDimension) => {
              return (
                <NavLink
                  to={`/org-dimensions/${orgDimension?.orgDimensionId}`}
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
                      <Card.Title>{orgDimension?.name}</Card.Title>
                      <Card.Text>
                        {orgDimension?.orgUnits?.items?.length} Org Units
                        <br />
                        {orgDimension?.orgUnitMemberships?.items?.length} Org
                        Accounts
                      </Card.Text>
                    </Card.Body>
                  </Card>
                </NavLink>
              );
            })}
            <br />
            <NewOrganizationalDimensionButton onCompleted={refetch} />
          </Nav>
        </Col>
        <Col md={"auto"}>
          <Outlet />
        </Col>
      </Row>
    </Container>
  );
};
