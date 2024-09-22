<template>
    <button class="button">
        <slot />
    </button>
</template>

<style scoped>
.button {
    background-color: #eeeeee;
    border-radius: 8px;
    border: 1px solid transparent;
    font-size: 0.9em;
    font-weight: 500;
    font-family: inherit;
    cursor: pointer;
    padding: 0.5em 1.2em;
    margin-left: 10px;
    transition-property: border-color;
    transition-duration: 0.25s;
}

.button:hover {
    border-color: #646cff;
}
</style>