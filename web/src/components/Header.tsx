import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import { NavLink } from "react-router-dom";

export function Header() {
  return (
    <Navbar bg="dark" expand="lg" sticky="top" variant="dark">
      <Container fluid>
        <Navbar.Brand>Fourcee</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto">
            <Nav.Link href="#organizations" as={NavLink} to="/org-structures">
              Organizations
            </Nav.Link>
            <Nav.Link href="#accounts" as={NavLink} to="/org-accounts">
              Accounts
            </Nav.Link>
            <Nav.Link href="#modules" as={NavLink} to="/module-groups">
              Modules
            </Nav.Link>
            <Nav.Link
              href="#propagations"
              as={NavLink}
              to="/module-propagations"
            >
              Propagations
            </Nav.Link>
            <Nav.Link href="#assignments" as={NavLink} to="/module-assignments">
              Assignments
            </Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}
