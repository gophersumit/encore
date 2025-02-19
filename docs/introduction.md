---
title: What is Encore?
subtitle: Escape complexity and put the fun back into backend development
---

Most developers can relate to the joyful creative feeling of programming. Many of us started programming at an early age, and the goal was to have fun. For a long time, we managed to hold on to this feeling even in a professional context, and we got to spend our days writing code to solve interesting business problems. However, in recent years the cloud has changed the lives of many backend developers. 

** Cloud services lets us build highly scalable systems, but there's a massive drawback. **

These days, most of our time is spent hand-crafting a sort of machine language for the cloud – _in the form of boilerplate_ – and on orchestrating and managing cloud services. This work is mind-numbingly repetitive, and doesn't add anything unique to the products we're building. What's worse is that the joyful feeling of programming is all but lost. This is what Encore is here to solve.

## An engine for backend, an epiphany for developers

For a long time, we've had to accept the status quo of boilerplate and cloud complexity. But it doesn't need to be like this. Encore is here to rock the boat and get backend developers focused on creative programming once again.

Taking inspiration from how game engines have empowered game developers to blast away with creative productivity, Encore gives you the same integrated special-purpose tooling for backend development. It does this by integrating the development process:

-   Write code with the Encore framework.
-   Run your local development environment with the Encore CLI.
-   Ship and run in production with the Encore Platform, connected to you own cloud.

<img src="/assets/docs/encore-way.png" title="The Encore Way" className="noshadow"/>

## A developer experience designed to help you stay in the flow

### 1. Develop with the Encore Framework

Setting up a productive development environment for building a modern backend application is very time-consuming. Many different services need to be integrated, complex configuration is required to connect all the pieces together, and there are many manual steps.

With Encore, you immediately get a very productive and fun developer experience when building backends and distributed systems:

-   Write regular Go functions and immediately get APIs.
-   Create multiple services simply by creating Go packages.
-   Get your local environment up and running in seconds with a single command: `encore run`
-   Backend primitives like databases, queues, and scheduled tasks, are native concepts that you express through Go code.
-   Best-practice solutions for common use cases like managing secrets, authentication, and observability are all built-in.
-   Local development dashboard with API explorer and distributed tracing.

### 2. Collaborate effortlessly with shorter feedback loops

Building advanced backend applications is often done in teams and there are great benefits that come with a collaborative development process. Encore makes it easy to collaborate faster with shorter feedback loops, thanks to its built-in collaboration features:

-   Automatically generate API documentation for all your APIs.
-   Integrate with GitHub to automatically set up each pull request as a preview environment.
-   Generate type-safe API frontend clients out of the box.
-   Co-develop a frontend and backend in realtime by exposing your local machine as a bespoke environment using `encore run --tunnel`.

### 3. Encore takes care of your entire DevOps process

With Encore you no longer need to spend your time on setting up complicated build and deployment pipelines, provisioning Kubernetes clusters and databases, or other busy-work. Instead, the Encore Platform takes care of setting it all up for you:

-   Fully fledged built-in CI/CD. Run `git push encore` to build, test, provision necessary infrastructure, and deploy.
-   Automatic provisioning of all your infrastructure for local, testing, and production environments in the cloud.
-   Free built-in hosting on Encore Cloud for development.
-   Connect your own cloud account with AWS/Azure/GCP to have Encore deploy there.
-   Create any number of environments and deploy to multiple clouds.

### 4. Debug production issues with ease

Encore comes with state of the art Distributed Tracing functionality, which dramatically simplifies finding the source of production issues. Is a service slow to respond? What data was being passed in to cause a particular bug to surface? Tracing makes it easy to find the answer.

## Backend use cases made simple with Encore

Encore is designed to help individual developers and teams be incredibly productive, and have more fun, when solving most backend use cases. There are many developers building with Encore, loving the experience when building things like:

-   CRUD backends and REST APIs.
-   Microservices backends for advanced web and mobile apps.
-   Highly performant APIs providing advanced business logic to 3rd parties.
-   And much more...

## Constraints that unlock great power

### So what's the catch?

The reason we normally spend so much time writing boilerplate, and configuring cloud services, is that the tools we rely on have no idea what we're trying to do. This means it's up to the developer to do almost all of the work. So in order to make a real improvement to the development process, we need a tool that does understand that. Encore was designed with this idea in mind, and it's deliberately opinionated in order to facilitate this understanding.

Just as constraints inspire creativity, we believe being opinionated means we can unlock incredible power for developers. So to enjoy all the Encore benefits, and the freedom to focus on creative programming instead of cloud complexity, you may need to let go of some personal preferences.

### Encore is built ground up for Go, and only Go

This is not to say Encore is built only for Go developers! We believe most backend developers will get incredible value from using Encore, and that learning Go should not stop anyone from adopting Encore.

Really, why is it that deciding on a programming language is often seen as the most important question when starting a new project? When you set out to build a new backend, there's often very few relevant and rational arguments for why one language is better than any other. The only real difference is personal taste.

### Congratulations, you don't have to decide

The same is true for dozens of other decisions we developers make when creating a backend; deciding how to structure the codebase, defining API schemas, etc, often come down to personal preference. This fragmentation creates a huge problem: When every stack looks different, all of our tools have to be general purpose.

When you adopt Encore many of these decisions are already made for you, in a way that ensures your application will follow modern best-practises. And when you run your application, Encore's open source parser and compiler checks that you're sticking to the framework. This means you're left to focus your energy on what really matters: writing your application's business logic.

### Meet the Encore Application Model

In order to deliver Encore's powerful features, Encore makes heavy use of static analysis — a fancy term for parsing and analyzing the code you write — to build up a very detailed understanding of how your backend works. This understanding is in the shape of a graph, and very closely represents your own mental model about your system: boxes and arrows, representing systems and services that communicate with other systems, pass data and connect to infrastructure. We call it the Encore Application Model.

Every Encore feature we build is designed with the Encore Application Model in mind. As developers we're not used to having tools that actually understand what we're doing. The height of intelligence has basically been "rename this function for me". No longer.

With the Encore Application Model we can build an entirely new generation of tools that understand what we as developers are trying to do, and then help us do it. And that's the key to developer productivity. And best of all? [It's open source](https://github.com/encoredev/encore).

<img src="/assets/docs/app-graph.png" title="Encore Application Model" className="noshadow mx-auto md:max-w-lg"/>

## One workflow from prototype to production

Many frameworks and tools designed for rapid development are great for prototyping, but become a liability as your application grows. *Encore is different.* It's designed from the ground up to be suitable for large scale software engineering. On top of this foundation are various features designed to make it incredibly easy to get started. This means you get the benefits of incredibly rapid prototyping, while your application gracefully handles increased requirements and scale as you move to production and beyond.

## Where to go from here

You made it to the end of the page – we hope you are as excited as we are about making backend development fun again!

Next, we recommend following the [Quick Start guide](/docs/quick-start). It gives you a taste of the Encore workflow and takes only around 5-10 minutes to complete, depending on your familiarity with Go.

After completing the guide, why not browse through the docs in your own areas of interest, or join [Slack](https://encore.dev/slack) to ask any questions you might have.

Finally, we recommend you follow and star the project on [GitHub](https://github.com/encoredev/encore) to stay up to date.
