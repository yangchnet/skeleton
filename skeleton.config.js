CWD = process.cwd()

module.exports = {
    apps: [
        {
            name: "openapi",
            script: "make run.openapi",
            cwd: CWD,
        },
        {
            name: "echo-service",
            script: "make run.echo-service",
            cwd: CWD,
            watch: [CWD + "/internal/echo"],
        },
        {
            name: "iam-service",
            script: "make run.iam-service",
            cwd: CWD,
            watch: [CWD + "/internal/iam"],
        }
    ]
}
