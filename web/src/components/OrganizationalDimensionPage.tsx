import React, { useState } from "react";
import {
  OrganizationalDimension,
  OrganizationalDimensions,
  OrganizationalUnit,
} from "../__generated__/graphql";
import { NavLink, useParams } from "react-router-dom";
import { useQuery, gql } from "@apollo/client";
import Table from "react-bootstrap/Table";
import Container from "react-bootstrap/Container";
import { Tree } from "react-organizational-chart";
import { TreeNode } from "react-organizational-chart";
import { NewOrganizationalUnitButton } from "./NewOrganizationalUnitButton";
import { GetOrgUnitTree, OrgUnitTreeNode } from "../utils/org_tree_rendering";

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
      modulePropagations {
        items {
          modulePropagationId
          moduleGroupId
          moduleVersionId
          orgUnitId
          orgDimensionId
          name
          description
        }
      }
      orgUnitMemberships {
        items {
          orgAccount {
            orgAccountId
            name
            cloudPlatform
            cloudIdentifier
          }
          orgUnit {
            orgUnitId
            name
          }
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

  const { loading, error, data } = useQuery<Response>(
    ORGANIZATIONAL_DIMENSION_QUERY,
    {
      variables: {
        orgDimensionId: organizationalDimensionId,
      },
      pollInterval: 1000,
    }
  );

  if (loading) return null;
  if (error) return <div>Error</div>;

  let orgUnitsMap = GetOrgUnitTree(
    organizationalDimensionId,
    data?.organizationalDimension.orgUnits.items ?? []
  );

  return (
    <Container style={{ minWidth: "50vw" }}>
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
      <NewOrganizationalUnitButton
        key={data?.organizationalDimension.orgDimensionId ?? ""}
        orgDimensionId={data?.organizationalDimension.orgDimensionId ?? ""}
        existingOrgUnits={data?.organizationalDimension.orgUnits.items ?? []}
      />
    </Container>
  );
};
