import { TreeNode } from "react-org-chart";
import { NavLink } from "react-router-dom";
import { Maybe, OrgUnit } from "../__generated__/graphql";

type OrgUnitTreeNode = {
  orgDimensionId: string;
  orgUnitId: string;
  name: string;
  children: OrgUnitTreeNode[];
};

interface OrgUnitTreeNodeProps {
  orgUnit: OrgUnitTreeNode | undefined;
}

export const OrgUnitTreeNode = (props: OrgUnitTreeNodeProps) => {
  return (
    <TreeNode
      label={
        <NavLink
          to={`/org-structures/${props.orgUnit?.orgDimensionId}/org-units/${props.orgUnit?.orgUnitId}`}
          style={({ isActive }) =>
            isActive
              ? {
                  color: "blue",
                  fontWeight: "bold",
                  padding: "8px",
                  borderRadius: "12px",
                  display: "inline-block",
                  border: "3px solid blue",
                }
              : {
                  color: "inherit",
                  textDecoration: "none",
                  padding: "8px",
                  borderRadius: "12px",
                  display: "inline-block",
                  border: "2px solid black",
                }
          }
        >
          {props.orgUnit?.name}
        </NavLink>
      }
    >
      {props.orgUnit?.children.map((child) => {
        return <OrgUnitTreeNode orgUnit={child} />;
      })}
    </TreeNode>
  );
};

export function GetOrgUnitTree(
  orgDimensionId: string,
  orgUnits: Maybe<OrgUnit>[]
): Map<string, OrgUnitTreeNode> {
  var rootOrgUnitId: string | null = null;
  let orgUnitsMap: Map<string, OrgUnitTreeNode> = new Map();
  for (let orgUnit of orgUnits) {
    if (orgUnit == null) continue;
    orgUnitsMap.set(orgUnit.id, {
      orgDimensionId: orgDimensionId ?? "",
      orgUnitId: orgUnit.id,
      name: orgUnit.name,
      children: [],
    });
  }

  for (let orgUnit of orgUnits) {
    if (orgUnit == null) continue;
    // add the org unit to it's parents children
    if (orgUnit.parentOrgUnitID) {
      let parentOrgUnit = orgUnitsMap.get(orgUnit.parentOrgUnitID);
      if (parentOrgUnit) {
        parentOrgUnit.children.push(orgUnitsMap.get(orgUnit.id)!);
      }
    }
  }

  return orgUnitsMap;
}
