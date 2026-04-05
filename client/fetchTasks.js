export default async function getTasks() {
    try {
        const tasksResp = await fetch("http://localhost:8081/tasks");
        return tasksResp.json();
    } catch (error) {
        console.error(error);
        return {};
    }
}