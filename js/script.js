// Copyright (c) 2023 Mohamad All rights reserved
//
// Created by: Mohamad
// Created on: May 8
// This file contains the JS functions for index.html

"use strict"

/*
 * This function checks the qualification of someone based on their age and the weekday
 */
function checkQualification() {
  // Input
  const weekday = document.getElementById("weekday").value
  const age = parseInt(document.getElementById("age").value)

  // Check if age is over 13
  if (age < 13) {
    document.getElementById("answer").innerHTML =
      "You are too young to go in the museum."
    // End program if age is under 13
    return
  }

  // Check if it is tuesday or thursday
  if (weekday == "tuesday" || weekday == "thursday") {
    document.getElementById("answer").innerHTML = "The museum is closed today."
  } else {
    document.getElementById("answer").innerHTML =
      "You can go to the museum today."
  }
}
