import React, { useState } from "react";
import { ModulePropagations } from "../__generated__/graphql";
import { NavLink, Outlet } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import Container from "react-bootstrap/Container";
import { ModulePropagationPage } from "./ModulePropagationPage";
import Table from "react-bootstrap/Table";
import { Card, Col, ListGroup, Nav, Row } from "react-bootstrap";
import { NewModulePropagationButton } from "./NewModulePropagationButton";
import { renderCloudPlatform } from "../utils/rendering";

const MODULE_PROPAGATIONS_QUERY = gql`
  query modulePropagations {
    modulePropagations(limit: 100) {
      items {
        modulePropagationId
        name
        moduleGroup {
          moduleGroupId
          name
          cloudPlatform
        }
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
  }
`;

type Response = {
  modulePropagations: ModulePropagations;
};

export const ModulePropagationsList = () => {
  const { loading, error, data, refetch } = useQuery<Response>(
    MODULE_PROPAGATIONS_QUERY,
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

            height: "calc(100vh - 3.5rem)",
            borderRight: "1px solid #dee2e6",
          }}
        >
          <h3>Module Propagations</h3>
          <Nav
            as={ListGroup}
            style={{
              maxHeight: "calc(100vh - 6rem)",
              flexDirection: "column",
              height: "100%",
              display: "flex",
              overflow: "auto",
              flexWrap: "nowrap",
            }}
          >
            {data?.modulePropagations.items.map((modulePropagation) => {
              return (
                <ListGroup.Item>
                  <Card>
                    <Card.Body>
                      <Card.Title>
                        <NavLink
                          to={`/module-propagations/${modulePropagation?.modulePropagationId}`}
                          style={({ isActive, isPending }) =>
                            isActive ? { color: "navy" } : {}
                          }
                        >
                          {modulePropagation?.name}
                        </NavLink>
                      </Card.Title>
                      <Card.Text>
                        <b>Module Group: </b>
                        {modulePropagation?.moduleGroup?.name}
                        <br />
                        <b>Module Version: </b>
                        {modulePropagation?.moduleVersion?.name}
                        <br />
                        <b>Org Unit: </b>
                        {modulePropagation?.orgDimension?.name}
                        {" / "}
                        {modulePropagation?.orgUnit?.name}
                      </Card.Text>
                    </Card.Body>
                  </Card>
                </ListGroup.Item>
              );
            })}
            <NewModulePropagationButton onCompleted={refetch} />
          </Nav>
        </Col>
        <Col md={"auto"}>
          <Outlet />
        </Col>
      </Row>
    </Container>
  );
};
