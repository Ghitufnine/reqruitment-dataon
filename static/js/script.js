// Function to toggle the visibility of child nodes
function toggleNode(element) {
    const childNodes = element.parentElement.querySelector("ul");
    if (childNodes) {
        if (childNodes.style.display === "none" || childNodes.style.display === "") {
            childNodes.style.display = "block";
            element.textContent = "[-]"; 
        } else {
            childNodes.style.display = "none";
            element.textContent = "[+]"; 
        }
    }
}

// Hide all child nodes initially when the document is loaded
document.addEventListener("DOMContentLoaded", () => {
    const childNodes = document.querySelectorAll("#department-tree ul ul"); // Select nested <ul> elements
    childNodes.forEach(node => {
        node.style.display = "none"; 
    });

    // Optionally, set toggle signs for the initially hidden nodes
    const toggleElements = document.querySelectorAll("#department-tree .toggle");
    toggleElements.forEach(element => {
        element.textContent = "[+]"; 
    });
});
