import React from "react";
import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import NavDropdown from "react-bootstrap/NavDropdown";
import { NavLink } from "react-router-dom";

export function Header() {
  return (
    <Navbar bg="light" expand="lg">
      <Container>
        <Navbar.Brand href="#organizations">TFOM</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link href="#organizations" as={NavLink} to="/org-dimensions">
              Organizations
            </Nav.Link>
            <Nav.Link href="#accounts" as={NavLink} to="/org-accounts">
              Accounts
            </Nav.Link>
            <Nav.Link href="#modules" as={NavLink} to="/modules">
              Modules
            </Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}
