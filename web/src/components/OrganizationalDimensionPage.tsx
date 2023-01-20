import React, { useState } from "react";
import {
  OrganizationalDimension,
  OrganizationalDimensions,
  OrganizationalUnit,
} from "../__generated__/graphql";
import { NavLink, Outlet, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import Container from "react-bootstrap/Container";
import { Tree } from "react-organizational-chart";
import { TreeNode } from "react-organizational-chart";
import { NewOrganizationalUnitButton } from "./NewOrganizationalUnitButton";
import { GetOrgUnitTree, OrgUnitTreeNode } from "../utils/org_tree_rendering";
import { Breadcrumb } from "react-bootstrap";

const ORGANIZATIONAL_DIMENSION_QUERY = gql`
  query organizationalDimension($orgDimensionId: ID!) {
    organizationalDimension(orgDimensionId: $orgDimensionId) {
      orgDimensionId
      name
      rootOrgUnitId
      orgUnits {
        items {
          orgUnitId
          name
          parentOrgUnitId
          hierarchy
        }
      }
    }
  }
`;

type Response = {
  organizationalDimension: OrganizationalDimension;
};

export const OrganizationalDimensionPage = () => {
  const params = useParams();

  const organizationalDimensionId = params.organizationalDimensionId
    ? params.organizationalDimensionId
    : "";

  const { loading, error, data, refetch } = useQuery<Response>(
    ORGANIZATIONAL_DIMENSION_QUERY,
    {
      variables: {
        orgDimensionId: organizationalDimensionId,
      },
      pollInterval: 5000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  let orgUnitsMap = GetOrgUnitTree(
    organizationalDimensionId,
    data?.organizationalDimension.orgUnits.items ?? []
  );

  return (
    <Container
      style={{
        paddingTop: "2rem",
        paddingBottom: "2rem",
      }}
      fluid
    >
      <Breadcrumb>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/" }}>
          Home
        </Breadcrumb.Item>
        <Breadcrumb.Item linkAs={NavLink} linkProps={{ to: "/org-dimensions" }}>
          Organizations
        </Breadcrumb.Item>
        <Breadcrumb.Item active>
          {data?.organizationalDimension.name} (
          {data?.organizationalDimension.orgDimensionId})
        </Breadcrumb.Item>
      </Breadcrumb>
      <Container style={{ minWidth: "50vw" }}>
        <NewOrganizationalUnitButton
          key={data?.organizationalDimension.orgDimensionId ?? ""}
          orgDimensionId={data?.organizationalDimension.orgDimensionId ?? ""}
          existingOrgUnits={data?.organizationalDimension.orgUnits.items ?? []}
          onCompleted={refetch}
        />
        <Tree
          label={<h1>{data?.organizationalDimension.name} Org Dimension</h1>}
          lineWidth={"2px"}
          nodePadding={"30px"}
        >
          <OrgUnitTreeNode
            orgUnit={orgUnitsMap.get(
              data?.organizationalDimension.rootOrgUnitId ?? ""
            )}
          />
        </Tree>
        <br />
      </Container>
      <Outlet />
    </Container>
  );
};
