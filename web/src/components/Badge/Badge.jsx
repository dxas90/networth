import React from "react";
import PropTypes from "prop-types";
import withStyles from "@material-ui/core/styles/withStyles";
import badgeStyle from "../../assets/jss/material-kit-react/components/badgeStyle.jsx";

function Badge({ ...props }) {
  const { classes, color, children } = props;
  return (
    <span className={classes.badge + " " + classes[color]}>{children}</span>
  );
}

Badge.defaultProps = {
  color: "gray"
};

Badge.propTypes = {
  classes: PropTypes.object.isRequired,
  children: PropTypes.string,
  color: PropTypes.oneOf([
    "primary",
    "warning",
    "danger",
    "success",
    "info",
    "rose",
    "gray"
  ])
};

export default withStyles(badgeStyle)(Badge);
