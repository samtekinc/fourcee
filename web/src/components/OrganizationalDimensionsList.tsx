import React, { useState } from "react";
import { OrganizationalDimensions } from "../__generated__/graphql";
import { NavLink } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import Container from "react-bootstrap/Container";
import { OrganizationalDimensionPage } from "./OrganizationalDimensionPage";
import Table from "react-bootstrap/Table";

const ORGANIZATIONAL_DIMENSIONS_QUERY = gql`
  query organizationalDimensions {
    organizationalDimensions(limit: 100) {
      items {
        orgDimensionId
        name
      }
    }
  }
`;

type Response = {
  organizationalDimensions: OrganizationalDimensions;
};

export const OrganizationalDimensionsList = () => {
  const [orgDimensionId, setOrgDimensionId] = useState("");

  const { loading, error, data } = useQuery<Response>(
    ORGANIZATIONAL_DIMENSIONS_QUERY,
    {}
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  return (
    <Container>
      <h1>Organizational Dimensions</h1>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>Org Dimension Id</th>
            <th>Name</th>
          </tr>
        </thead>
        <tbody>
          {data?.organizationalDimensions.items.map((dimension) => {
            return (
              <tr>
                <td>
                  <NavLink to={`/org-dimensions/${dimension?.orgDimensionId}`}>
                    {dimension?.orgDimensionId}
                  </NavLink>
                </td>
                <td>{dimension?.name}</td>
              </tr>
            );
          })}
        </tbody>
      </Table>
    </Container>
  );
};
